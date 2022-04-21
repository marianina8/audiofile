package volumedetect

import (
	"audiofile/models"
	"bufio"
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func Extract(m *models.Audio) error {
	cmdArgs := []string{
		"-i", m.Path,
		"-af", "volumedetect",
		"-f", "null", "/dev/null",
	}

	out, err := exec.Command("ffmpeg", cmdArgs...).CombinedOutput()
	if err != nil {
		return err
	}

	var volume models.VolumeDetect
	err = parseOutput(out, &volume)
	if err != nil {
		return err
	}
	m.Metadata.VolumeDetect = volume
	return nil
}

func parseOutput(output []byte, model *models.VolumeDetect) error {
	// Sample output to parse
	//
	// [Parsed_volumedetect_0 @ 0x7f98b07009c0] n_samples: 183939
	// [Parsed_volumedetect_0 @ 0x7f98b07009c0] mean_volume: -36.5 dB
	// [Parsed_volumedetect_0 @ 0x7f98b07009c0] max_volume: -15.3 dB
	// [Parsed_volumedetect_0 @ 0x7f98b07009c0] histogram_15db: 9
	// [Parsed_volumedetect_0 @ 0x7f98b07009c0] histogram_16db: 8
	// [Parsed_volumedetect_0 @ 0x7f98b07009c0] histogram_17db: 16
	// [Parsed_volumedetect_0 @ 0x7f98b07009c0] histogram_18db: 41
	// [Parsed_volumedetect_0 @ 0x7f98b07009c0] histogram_19db: 71
	// [Parsed_volumedetect_0 @ 0x7f98b07009c0] histogram_20db: 90

	scanner := bufio.NewScanner(bytes.NewReader(output))
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[Parsed_volumedetect") {

			// cleanup of excessive whitespace and splitting results
			chunk := strings.Split(line, "]")
			chunk[1] = strings.TrimSpace(chunk[1])
			data := strings.Split(chunk[1], ":")
			data[1] = strings.TrimSpace(data[1])
			switch data[0] {
			case "mean_volume":
				model.MeanVolumeDB = parseVolumeDB(data[1])
			case "max_volume":
				model.MaxVolumeDB = parseVolumeDB(data[1])
			case "n_samples":
				model.NSamples = parseInt(data[1])
			case "histogram_0db":
				model.Histogram0DB = parseInt(data[1])
			case "histogram_1db":
				model.Histogram1DB = parseInt(data[1])
			case "histogram_2db":
				model.Histogram2DB = parseInt(data[1])
			case "histogram_3db":
				model.Histogram3DB = parseInt(data[1])
			case "histogram_4db":
				model.Histogram4DB = parseInt(data[1])
			case "histogram_5db":
				model.Histogram5DB = parseInt(data[1])
			case "histogram_6db":
				model.Histogram6DB = parseInt(data[1])
			case "histogram_7db":
				model.Histogram7DB = parseInt(data[1])
			case "histogram_8db":
				model.Histogram8DB = parseInt(data[1])

			// channel 2
			case "histogram_15db":
				model.Histogram15DB = parseInt(data[1])
			case "histogram_16db":
				model.Histogram16DB = parseInt(data[1])
			case "histogram_17db":
				model.Histogram17DB = parseInt(data[1])
			case "histogram_18db":
				model.Histogram18DB = parseInt(data[1])
			case "histogram_19db":
				model.Histogram19DB = parseInt(data[1])
			case "histogram_20db":
				model.Histogram20DB = parseInt(data[1])
			}
		}
	}
	err := scanner.Err()
	if err != nil {
		return err
	}
	return nil
}

func parseInt(value string) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return v
}

func parseVolumeDB(value string) float64 {
	parts := strings.Fields(value)
	if len(parts) != 2 {
		return 0
	}
	if parts[1] != "dB" {
		return 0
	}

	v, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0
	}
	return v
}
