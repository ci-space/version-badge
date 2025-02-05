package generator

import (
	"errors"
	"github.com/ci-space/version-badge/fonts"
)

type Font string

const (
	FontUnspecified Font = ""
	FontVerdana     Font = "verdana"
	FontArial       Font = "arial"
	FontComicSansMS Font = "comic-sans-ms"
)

func (f Font) Valid() bool {
	switch f {
	case FontVerdana:
		return true
	case FontArial:
		return true
	case FontComicSansMS:
		return true
		return true
	case FontUnspecified:
		return false
	}

	return false
}

func resolveTTF(font Font) ([]byte, error) {
	switch font {
	case FontUnspecified:
		return nil, errors.New("font unspecified")
	case FontVerdana:
		return fonts.Verdana, nil
	case FontArial:
		return fonts.Arial, nil
	case FontComicSansMS:
		return fonts.ComicSansMS, nil
	}

	return nil, errors.New("font unspecified")
}
