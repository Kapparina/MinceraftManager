package installation

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Kapparina/MinceraftManager/backend/pkg/minecraft"
)

type Installation struct {
	Id          string    `json:"id"`
	Type        string    `json:"type"`
	MainClass   string    `json:"mainClass"`
	ReleaseTime time.Time `json:"releaseTime"`
	Time        time.Time `json:"time"`
}

func GetGameVersions() ([]Installation, error) {
	installationsDir := filepath.Join(minecraft.InstallationDir, "versions")
	entries, err := os.ReadDir(installationsDir)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	results := make(chan Installation, len(entries))
	errs := make(chan error, len(entries))
	var installations []Installation

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		wg.Add(1)
		go func(dirName string) {
			defer wg.Done()
			data, err := os.ReadFile(filepath.Join(installationsDir, dirName, dirName+".json"))
			if err != nil {
				errs <- err
				return
			}
			var installation Installation
			if err = json.Unmarshal(data, &installation); err != nil {
				errs <- err
				return
			}
			results <- installation
		}(entry.Name())
	}
	go func() {
		wg.Wait()
		close(results)
		close(errs)
	}()
	for i := range results {
		installations = append(installations, i)
	}
	var bigErr error
	for e := range errs {
		errors.Join(bigErr, e)
	}
	return installations, bigErr
}
