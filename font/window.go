//go:build windows
// +build windows

package font

import (
	"os"
	"path/filepath"
)

func FontDirectory() []string {
	return []string{
		filepath.Join(os.ExpandEnv("${windir}"), "Fonts"),
		filepath.Join(os.ExpandEnv("${localappdata}"), "Microsoft", "Windows", "Fonts"),
	}
}
