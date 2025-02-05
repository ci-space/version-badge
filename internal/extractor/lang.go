package extractor

import (
	"errors"
	"github.com/artarts36/depexplorer"
	"github.com/ci-space/version-badge/internal/version"
)

func Language(depFile *depexplorer.File, _ string) (*version.Version, error) {
	if depFile.Language.Version == nil {
		return nil, errors.New("file not contains language version")
	}

	return &version.Version{
		Object:  string(depFile.Language.Name),
		Version: depFile.Language.Version.Full,
	}, nil
}
