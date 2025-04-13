package minecraft

import (
	"os"
	"path/filepath"
)

var (
	InstallationDir string
	ModDir          string
	ResourcePackDir string
	ShaderDir       string
)

func init() {
	var err error
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	InstallationDir = filepath.Join(userDir, "AppData/Roaming/.minecraft")
	ModDir = filepath.Join(InstallationDir, "mods")
	ResourcePackDir = filepath.Join(InstallationDir, "resourcepacks")
	ShaderDir = filepath.Join(InstallationDir, "shaderpacks")
}
