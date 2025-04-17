package pack

type McMetaFile struct {
	Pack struct {
		PackFormat       int    `json:"pack_format"`
		Description      string `json:"description"`
		SupportedFormats struct {
			MinInclusive int `json:"min_inclusive"`
			MaxInclusive int `json:"max_inclusive"`
		} `json:"supported_formats"`
	} `json:"pack"`
	Overlays struct {
		Entries []struct {
			Directory string `json:"directory"`
			Formats   []int  `json:"formats"`
		} `json:"entries"`
	} `json:"overlays"`
}
