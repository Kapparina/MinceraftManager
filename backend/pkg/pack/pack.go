package pack

import (
	"archive/zip"
	"encoding/json"
	"io"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
)

type Pack interface {
	PackDescription() string
}

func ReadPack(path string) (Pack, error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer func(reader *zip.ReadCloser) {
		err = reader.Close()
	}(reader)
	fileMap := make(map[string]*zip.File)
	var mcMetaData json.RawMessage

	for _, f := range reader.File {
		fileMap[f.Name] = f
		slog.Debug(f.Name)
		if f.FileInfo().Name() == "pack.mcmeta" {
			if c, readErr := f.Open(); readErr == nil {
				contents := make([]byte, f.FileInfo().Size())
				if contents, readErr = io.ReadAll(c); err != nil {
					return nil, readErr
				}
				mcMetaData = contents
			}
		}
	}
	switch filepath.Ext(path) {
	case ".zip":
		var pack ResourcePack
		if err = json.Unmarshal(mcMetaData, &pack.McMetaFile); err != nil {
			return nil, err
		}
		return &pack, nil
	case ".jar":
		for n, f := range fileMap {
			if strings.HasPrefix(n, "META-INF/") && strings.HasSuffix(n, ".mods.toml") {
				var pack ModPack
				if err = json.Unmarshal(mcMetaData, &pack.McMetaFile); err != nil {
					return nil, err
				}
				c, readErr := f.Open()
				if readErr != nil {
					return nil, readErr
				}
				contents := make([]byte, f.FileInfo().Size())
				if contents, readErr = io.ReadAll(c); readErr != nil {
					return nil, readErr
				}
				if err = toml.Unmarshal(contents, &pack.ModsToml); err != nil {
					return nil, err
				}
				return &pack, nil
			}
		}
	default:
		return nil, errors.New("unsupported pack type")
	}
	return nil, errors.New("pack not found")
}
