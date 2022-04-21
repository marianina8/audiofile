package audiopeak

import (
	"audiofile/models"
	"bufio"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

func Extract(m *models.Audio) error {
	cmdArgs := []string{
		"-nostats",
		"-i", m.Path,
		"-filter_complex", "ebur128=peak=true",
		"-f", "null", "-",
	}

	b, err := exec.Command("ffmpeg", cmdArgs...).CombinedOutput()
	if err != nil {
		return nil
	}

	split := strings.Split(string(b), "Summary:")
	out := split[len(split)-1]
	out = strings.TrimSpace(out)

	var audio models.AudioPeak
	if err = parseOutput(out, &audio); err != nil {
		return err
	}
	m.Metadata.Audiopeak = audio
	return nil
}

func parseOutput(ffmpegOutput string, audio *models.AudioPeak) error {
	scanner := bufio.NewScanner(strings.NewReader(ffmpegOutput))
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "I:") {
			res := strings.Split(line, ":")
			res[1] = strings.TrimSpace(res[1])
			audio.IntegratedLoudness.ILUFs = parseUnits(res[1], "LUFS")
		}

		if strings.Contains(line, "Threshold:") {
			res := strings.Split(line, ":")
			res[1] = strings.TrimSpace(res[1])
			if audio.IntegratedLoudness.ThresholdLUFS == nil {
				audio.IntegratedLoudness.ThresholdLUFS = parseUnits(res[1], "LUFS")
			} else {
				audio.LoudnessRange.ThresholdLUFS = parseUnits(res[1], "LUFS")
			}
		}

		if strings.Contains(line, "LRA:") {
			res := strings.Split(line, ":")
			res[1] = strings.TrimSpace(res[1])
			audio.LoudnessRange.LraLU = parseUnits(res[1], "LU")
		}

		if strings.Contains(line, "LRA low:") {
			res := strings.Split(line, ":")
			res[1] = strings.TrimSpace(res[1])
			audio.LoudnessRange.LraLowLUFS = parseUnits(res[1], "LUFS")
		}

		if strings.Contains(line, "LRA high:") {
			res := strings.Split(line, ":")
			res[1] = strings.TrimSpace(res[1])
			audio.LoudnessRange.LraHighLUFS = parseUnits(res[1], "LUFS")
		}

		if strings.Contains(line, "Peak:") {
			res := strings.Split(line, ":")
			res[1] = strings.TrimSpace(res[1])
			audio.TruePeakdBFS = parseUnits(res[1], "dBFS")
		}
	}

	return nil
}

func parseUnits(value string, units string) *float64 {
	parts := strings.Fields(value)
	if len(parts) != 2 {
		return nil
	}

	if parts[1] != units {
		return nil
	}

	v, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil
	}

	if v == math.Inf(-1) {
		return nil
	}

	return &v
}
