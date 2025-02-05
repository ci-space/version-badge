package main

import (
	"context"
	"fmt"
	"github.com/ci-space/version-badge/internal/generator"
	"strconv"

	cli "github.com/artarts36/singlecli"

	"github.com/ci-space/version-badge/internal/cmd"
)

func main() {
	app := &cli.App{
		BuildInfo: &cli.BuildInfo{
			Name:        "version-badge",
			Description: "Console app and GitHub Action for generate SVG badges with version of language or dependency",
			Version:     "v0.1.0",
		},
		Action: run,
		Args: []*cli.ArgDefinition{
			{
				Name:        "out",
				Required:    true,
				Description: "path to output file",
			},
		},
		Opts: []*cli.OptDefinition{
			{
				Name:        "from",
				WithValue:   true,
				Description: "Path to dependencies file (go.mod, composer.json, package.json)",
			},
			{
				Name:        "object",
				WithValue:   true,
				Description: "Dependency name",
			},
			{
				Name:      "style",
				WithValue: true,
			},
			{
				Name:        "font",
				WithValue:   true,
				Description: "Default: Verdana",
			},
			{
				Name:        "font-size",
				WithValue:   true,
				Description: "Default: 11",
			},
			{
				Name: "short-object-name",
			},
			{
				Name:        "color",
				WithValue:   true,
				Description: "Background color for right side. HEX or blue/green/brightgreen",
			},
		},
	}

	app.RunWithGlobalArgs(context.Background())
}

func run(ctx *cli.Context) error {
	params := cmd.Params{
		Object:    ctx.Opts["object"],
		From:      ctx.Opts["from"],
		Font:      ctx.Opts["font"],
		Out:       ctx.Args["out"],
		ShortName: ctx.HasOpt("short-object-name"),
		Color:     ctx.Opts["color"],
	}

	if fontSize, ok := ctx.Opts["font-size"]; ok {
		fSize, err := strconv.Atoi(fontSize)
		if err != nil {
			return fmt.Errorf("failed to convert font size to int: %w", err)
		}

		params.FontSize = fSize
	}

	command := cmd.NewCommand(generator.NewGenerator())

	return command.Run(params)
}
