package extractor

import (
	"fmt"

	"github.com/artarts36/depexplorer"

	"github.com/ci-space/version-badge/internal/version"
)

func Dependency(depFile *depexplorer.File, depName string) (*version.Version, error) {
	var dep *depexplorer.Dependency

	for _, dependency := range depFile.Dependencies {
		if dependency.Name == depName {
			dep = dependency
			break
		}
	}

	if dep == nil {
		return nil, fmt.Errorf("dependency %q not found", depName)
	}

	return &version.Version{
		Object:  dep.Name,
		Version: dep.Version.Full,
	}, nil
}
