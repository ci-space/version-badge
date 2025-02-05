package generator

import (
	"errors"
	"fmt"

	"github.com/essentialkaos/go-badge"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(bg Badge) ([]byte, error) {
	ttf, err := resolveTTF(bg.Font)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve ttf: %w", err)
	}

	generator, err := badge.NewGeneratorFromBytes(ttf, bg.FontSize)
	if err != nil {
		return nil, fmt.Errorf("unable to create badge generator: %w", err)
	}

	return g.gen(generator, bg)
}

func (g *Generator) gen(generator *badge.Generator, bg Badge) ([]byte, error) {
	switch bg.Style {
	case StyleFlat:
		return generator.GenerateFlat(bg.Object, bg.Version, bg.Color), nil
	case StylePlastic:
		return generator.GeneratePlastic(bg.Object, bg.Version, bg.Color), nil
	case StyleFlagSquare:
		return generator.GenerateFlatSquare(bg.Object, bg.Version, bg.Color), nil
	case StyleUnspecified:
		return nil, errors.New("badge style not specified")
	default:
		return nil, errors.New("badge style not specified")
	}
}
