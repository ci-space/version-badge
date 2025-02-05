package depfile

import (
	"fmt"
	"os"

	"github.com/artarts36/depexplorer"
)

func Fetch(path string) (*depexplorer.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	fStat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to stat file: %w", err)
	}

	if fStat.IsDir() {
		files, serr := depexplorer.ScanProjectDir(path)
		if serr != nil {
			return nil, fmt.Errorf("failed to scan dir: %w", serr)
		}

		if len(files) > 1 {
			return nil, fmt.Errorf("directory contains multiple dependency files: %w. You should specify path", err)
		}

		for _, f := range files {
			return f, nil
		}
	}

	return depexplorer.Guess(path)
}
