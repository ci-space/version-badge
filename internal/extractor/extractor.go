package extractor

import (
	"github.com/artarts36/depexplorer"

	"github.com/ci-space/version-badge/internal/version"
)

type Extractor func(file *depexplorer.File, depName string) (*version.Version, error)
