//go:build darwin
// +build darwin

package font

import (
	"os"
	"path/filepath"
)

func FontDirectory() []string {
	homeDir, _ := os.UserHomeDir()

	return []string{
		filepath.Join(homeDir, "Library/Fonts"),
		"/Library/Fonts/",
		"/System/Library/Fonts/"
	}
}
