package models

type Metadata struct {
	Audiopeak    AudioPeak    `json:"audiopeak"`
	VolumeDetect VolumeDetect `json:"volumedetect"`
	Tags         Tags         `json:"tags"`
	Transcript   string       `json:"transcript"`
}
