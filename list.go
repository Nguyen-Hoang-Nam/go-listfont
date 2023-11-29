package listfont

import (
	"io/fs"
	"path"
	"path/filepath"

	"github.com/Nguyen-Hoang-Nam/go-listfont/cache"
	"github.com/Nguyen-Hoang-Nam/go-listfont/font"
)

var supportedExts = []string{".ttf", ".otf", ".ttc"}

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func UpdateCache() error {
	fonts := list()
	return cache.Write(fonts)
}

func List() []font.Font {
	fonts, err := cache.Read()
	if err != nil {
		fonts := list()
		cache.Write(fonts)

		return fonts
	}

	return fonts

}

func list() []font.Font {
	fonts := []font.Font{}
	dirs := font.FontDirectory()

	for _, dir := range dirs {
		filepath.WalkDir(dir, func(s string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() {
				ext := path.Ext(s)
				if contains(supportedExts, ext) {
					fs, err := font.Parse(s)
					if err != nil {
						return nil
					}

					fonts = append(fonts, fs...)
				}
			}

			return nil
		})
	}

	return fonts
}
