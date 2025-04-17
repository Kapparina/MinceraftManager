package pack

import (
	"testing"
)

func TestReadPack(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		expected  string
		shouldErr bool
	}{
		{
			name:     "valid zip with resource pack",
			path:     "./testdata/resource-pack.zip",
			expected: "§5Alacrity §232x pack \n§7Author: §3Satellence§8",
		},
		{
			name:     "valid jar with mod pack",
			path:     "./testdata/mod.jar",
			expected: "Jade Resources",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pack, err := ReadPack(tt.path)
			if tt.shouldErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
					return
				}
				if pack.PackDescription() != tt.expected {
					t.Errorf("expected %s, got %s", tt.expected, pack.PackDescription())
				}
			}
		})
	}
}
