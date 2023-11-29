//go:build linux
// +build linux

package font

import (
	"os"
	"path/filepath"
)

func FontDirectory() []string {
	homeDir, _ := os.UserHomeDir()

	paths := []string{filepath.Join(homeDir, ".fonts")}

	if xdg_data_home := os.Getenv("XDG_DATA_HOME"); xdg_data_home != "" {
		paths = append(paths, filepath.Join(xdg_data_home, "fonts"))
	} else {
		paths = append(paths, filepath.Join(homeDir, ".local/share/fonts"))
	}

	if xdg_data_dirs := os.Getenv("XDG_DATA_DIRS"); xdg_data_dirs != "" {
		for _, xdg_data_dir := range filepath.SplitList(xdg_data_dirs) {
			paths = append(paths, filepath.Join(xdg_data_dir, "fonts"))
		}
		return paths
	} else {
		paths = append(paths, "/usr/local/share/fonts", "/usr/share/fonts")
	}

	return paths
}
