package cmd

import (
	"fmt"
	"github.com/artarts36/depexplorer"
	"github.com/ci-space/version-badge/internal/depfile"
	"github.com/ci-space/version-badge/internal/extractor"
	"github.com/ci-space/version-badge/internal/generator"
	"github.com/ci-space/version-badge/internal/version"
	"os"
)

const (
	defaultFontSize = 11
)

type Command struct {
	gen *generator.Generator
}

type Params struct {
	Object    string
	ShortName bool
	From      string

	Style    string
	Font     string
	FontSize int

	Out string
}

func NewCommand(gen *generator.Generator) *Command {
	return &Command{
		gen: gen,
	}
}

func (c *Command) Run(params Params) error {
	bg := c.makeBadge(params)

	path := params.From
	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}
	}

	depFile, err := depfile.Fetch(path)
	if err != nil {
		return fmt.Errorf("failed to fetch dependency file: %w", err)
	}

	vers, err := c.extractVersion(depFile, params)
	if err != nil {
		return fmt.Errorf("failed to extract version: %w", err)
	}

	bg.Object = vers.Object
	bg.Version = vers.Version

	if params.ShortName {
		bg.Object = vers.ShortObjectName()
	}

	return c.generate(bg, params.Out)
}

func (c *Command) makeBadge(params Params) generator.Badge {
	bg := generator.Badge{
		Style:    generator.Style(params.Style),
		Font:     generator.Font(params.Font),
		FontSize: params.FontSize,
		Color:    "#97ca00",
	}

	if !bg.Style.Valid() {
		bg.Style = generator.StyleFlat
	}
	if !bg.Font.Valid() {
		bg.Font = generator.FontVerdana
	}
	if bg.FontSize == 0 {
		bg.FontSize = defaultFontSize
	}

	return bg
}

func (c *Command) generate(bg generator.Badge, out string) error {
	content, err := c.gen.Generate(bg)
	if err != nil {
		return err
	}

	return os.WriteFile(out, content, 0755)
}

func (c *Command) extractVersion(depFile *depexplorer.File, params Params) (*version.Version, error) {
	extract := extractor.Language
	if params.Object != "" {
		extract = extractor.Dependency
	}

	return extract(depFile, params.Object)
}
