package installation

import "time"

type Installation struct {
	Id          string    `json:"id"`
	Type        string    `json:"type"`
	MainClass   string    `json:"mainClass"`
	ReleaseTime time.Time `json:"releaseTime"`
	Time        time.Time `json:"time"`
}
