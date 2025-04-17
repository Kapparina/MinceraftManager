package pack

//goland:noinspection GoNameStartsWithPackageName
type ResourcePack struct {
	McMetaFile `json:"pack"`
	Languages  []string `json:"language"`
}

func (r *ResourcePack) PackDescription() string {
	return r.Pack.Description
}
