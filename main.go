package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type options struct {
	Image   string
	Build   string
	Embed   bool
	Workdir string
	Env     []string
	Volumes []string
}

func main() {
	app := &cli.App{
		Name:   "binny",
		Usage:  "create an executable from a docker image",
		Action: generate,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Required: true,
				Name:     "image",
				Usage:    "name of your docker image",
			},
			&cli.BoolFlag{
				Name:  "embed",
				Usage: "embed a docker image in the binary",
			},
			&cli.StringFlag{
				Name:  "build",
				Usage: "build a docker image upon install",
			},
			&cli.StringFlag{
				Name:    "workdir",
				Aliases: []string{"w"},
				Usage:   "mount the user's current directory in the image",
			},
			&cli.StringSliceFlag{
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "whitelist environment variables",
			},
			&cli.StringSliceFlag{
				Name:    "volume",
				Aliases: []string{"v"},
				Usage:   "bind mount a volume",
			},
		},
	}

	app.Run(os.Args)
}

func generate(c *cli.Context) error {
	opts := parseOptions(c)
	fmt.Printf("%+v\n", opts)
	return nil
}

func parseOptions(c *cli.Context) options {
	return options{
		Image:   c.String("image"),
		Build:   c.String("build"),
		Embed:   c.Bool("embed"),
		Workdir: c.String("workdir"),
		Env:     c.StringSlice("env"),
		Volumes: c.StringSlice("volume"),
	}
}
