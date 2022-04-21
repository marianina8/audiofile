package models

// AudioPeak represents technical metadata extracted by the audiopeak extractor
type AudioPeak struct {
	IntegratedLoudness struct {
		ILUFs         *float64 `json:"i_lufs"`
		ThresholdLUFS *float64 `json:"threshold_lufs"`
	} `json:"integrated_loudness"`
	LoudnessRange struct {
		LraLU         *float64 `json:"lra_lu"`
		ThresholdLUFS *float64 `json:"threshold_lufs"`
		LraLowLUFS    *float64 `json:"lra_low_lufs"`
		LraHighLUFS   *float64 `json:"lra_high_lufs"`
	} `json:"loudness_range"`
	TruePeakdBFS *float64 `json:"true_peak_dbfs"`
}
