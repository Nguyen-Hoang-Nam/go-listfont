package font

import (
	"errors"
	"os"
	"path"

	"golang.org/x/image/font/sfnt"
)

type Font struct {
	Uri       string
	Family    string
	SubFamily string
}

func convertFont(f *sfnt.Font) map[string]string {
	nameIDs := []sfnt.NameID{
		sfnt.NameIDFamily,
		sfnt.NameIDSubfamily,
	}

	labels := []string{
		"NameIDFamily",
		"NameIDSubfamily",
	}

	NameIDInfo := map[string]string{}
	for i, nameID := range nameIDs {
		label := labels[i]

		value, err := f.Name(nil, nameID)
		if err != nil {
			NameIDInfo[label] = ""
			continue
		}

		NameIDInfo[label] = value
	}

	return NameIDInfo
}

func parseFont(s string, b []byte) (fonts []Font, err error) {
	f, err := sfnt.Parse(b)
	if err != nil {
		return nil, err
	}

	NameIDInfo := convertFont(f)

	return []Font{{Uri: s, Family: NameIDInfo["NameIDFamily"], SubFamily: NameIDInfo["NameIDSubfamily"]}}, err
}

func parseCollection(s string, b []byte) ([]Font, error) {
	c, err := sfnt.ParseCollection(b)
	if err != nil {
		return nil, err
	}

	fonts := []Font{}

	numPoints := c.NumFonts()
	for i := 1; i < numPoints; i++ {
		f, err := c.Font(i)
		if err != nil {
			return nil, err
		}

		NameIDInfo := convertFont(f)
		fonts = append(fonts, Font{Uri: s, Family: NameIDInfo["NameIDFamily"], SubFamily: NameIDInfo["NameIDSubfamily"]})
	}

	return fonts, err
}

func Parse(filename string) ([]Font, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	switch path.Ext(filename) {
	case ".ttc":
		return parseCollection(filename, b)
	case ".ttf", ".otf":
		return parseFont(filename, b)
	default:
		return nil, errors.New("Font format not supported")
	}
}
