package minecraft

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	InstallationDir string
	ModDir          string
	ResourcePackDir string
	ShaderDir       string
)

func init() {
	var err error
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	switch runtime.GOOS {
	case "windows":
		InstallationDir = filepath.Join(homeDir, "AppData/Roaming/.minecraft")
	case "darwin":
		InstallationDir = filepath.Join(homeDir, "Library/Application Support/minecraft")
	case "linux":
		InstallationDir = filepath.Join(homeDir, ".minecraft")
	}
	ModDir = filepath.Join(InstallationDir, "mods")
	ResourcePackDir = filepath.Join(InstallationDir, "resourcepacks")
	ShaderDir = filepath.Join(InstallationDir, "shaderpacks")
}
