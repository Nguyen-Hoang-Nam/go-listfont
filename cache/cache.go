package cache

import (
	"bytes"
	"encoding/gob"
	"os"
	"path"

	"github.com/Nguyen-Hoang-Nam/go-listfont/font"
)

const (
	CACHE_DIRECTORY = "GO_LISTFONT_CACHE_DIRECTORY"
	CACHE_FILE_NAME = "fonts.bin"
)

func cacheDir() (string, error) {
	if value := os.Getenv(CACHE_DIRECTORY); value != "" {
		return value, nil
	}

	configDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return path.Join(configDir, "go-listfont"), nil
}

func cachePath() (string, error) {
	cacheDir, err := cacheDir()
	if err != nil {
		return "", err
	}

	return path.Join(cacheDir, CACHE_FILE_NAME), nil
}

func Write(fonts []font.Font) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(fonts)

	p, err := cachePath()
	if err != nil {
		return err
	}

	if _, err = os.Stat(p); os.IsNotExist(err) {
		dir, _ := cacheDir()
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	err = os.WriteFile(p, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func Read() ([]font.Font, error) {
	p, err := cachePath()
	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(p); os.IsNotExist(err) {
		return nil, err
	}

	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	var fonts []font.Font

	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&fonts)
	if err != nil {
		return nil, err
	}

	return fonts, nil
}
