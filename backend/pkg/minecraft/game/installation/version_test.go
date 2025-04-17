package installation

import (
	"testing"
)

func TestGetGameVersions(t *testing.T) {
	tests := []struct {
		name        string
		expectErr   bool
		minExpected int
	}{
		{
			name:        "read real installations",
			expectErr:   false,
			minExpected: 0, // Assuming there might be some installations but it's okay if none
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			installations, err := GetGameVersions()

			if tt.expectErr && err == nil {
				t.Fatalf("expected an error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("did not expect an error but got: %v", err)
			}
			if len(installations) < tt.minExpected {
				t.Fatalf("expected at least %d installations, but got %d", tt.minExpected, len(installations))
			}
			t.Logf("got %d installations", len(installations))
			for _, installation := range installations {
				t.Logf("%+v", installation)
			}
		})
	}
}
