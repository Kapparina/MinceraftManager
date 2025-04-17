package pack

type ModPack struct {
	McMetaFile
	ModsToml
}

func (m *ModPack) PackDescription() string {
	return m.Pack.Description
}
