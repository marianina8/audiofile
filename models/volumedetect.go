package models

// VolumeDetect describes volume data
type VolumeDetect struct {
	NSamples     int     `json:"n_samples,omitempty"`
	MeanVolumeDB float64 `json:"mean_volume_db,omitempty"`
	MaxVolumeDB  float64 `json:"max_volume_db,omitempty"`
	Histogram0DB int     `json:"histogram_0_db,omitempty"`
	Histogram1DB int     `json:"histogram_1_db,omitempty"`
	Histogram2DB int     `json:"histogram_2_db,omitempty"`
	Histogram3DB int     `json:"histogram_3_db,omitempty"`
	Histogram4DB int     `json:"histogram_4_db,omitempty"`
	Histogram5DB int     `json:"histogram_5_db,omitempty"`
	Histogram6DB int     `json:"histogram_6_db,omitempty"`
	Histogram7DB int     `json:"histogram_7_db,omitempty"`
	Histogram8DB int     `json:"histogram_8_db,omitempty"`

	// 2 channel
	Histogram15DB int `json:"histogram_15_db,omitempty"`
	Histogram16DB int `json:"histogram_16_db,omitempty"`
	Histogram17DB int `json:"histogram_17_db,omitempty"`
	Histogram18DB int `json:"histogram_18_db,omitempty"`
	Histogram19DB int `json:"histogram_19_db,omitempty"`
	Histogram20DB int `json:"histogram_20)db,omitempty"`
}
