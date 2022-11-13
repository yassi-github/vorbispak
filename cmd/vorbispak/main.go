package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/yassi-github/vorbispak"
)

func main() {
	var out string
	var name string
	var pname string
	var ifile string

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to get current working directory")
	}

	app := &cli.App{
		Usage: "cut or merge ogg vorbis pak file",
		Commands: []*cli.Command{
			{
				Name:  "cut",
				Usage: "cut pak file to ogg vorbis files",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "out",
						Aliases:     []string{"o"},
						Usage:       "output ogg file `path`",
						Destination: &out,
						Value:       cwd,
					},
					&cli.StringFlag{
						Name:        "prefix-name",
						Aliases:     []string{"p"},
						Usage:       "output ogg file name `prefix`",
						Destination: &pname,
					},
					&cli.StringFlag{
						Name:        "index-file",
						Aliases:     []string{"i"},
						Usage:       "index file `path` to input",
						Destination: &ifile,
					},
				},
				Action: func(cCtx *cli.Context) error {
					var pak_file string
					if cCtx.NArg() > 0 {
						pak_file = cCtx.Args().First()
					} else {
						return cli.Exit("Argument `pak-file` required.", 1)
					}

					vorbispak.Cut(out, pname, ifile, pak_file)
					return nil
				},
			},
			{
				Name:  "merge",
				Usage: "merge ogg vorbis files to pak file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "out",
						Aliases:     []string{"o"},
						Usage:       "output pak file `path`",
						Destination: &out,
						Value:       cwd,
					},
					&cli.StringFlag{
						Name:        "name",
						Aliases:     []string{"n"},
						Usage:       "output pak file `name`",
						Destination: &name,
					},
					&cli.StringFlag{
						Name:        "index-file",
						Aliases:     []string{"i"},
						Usage:       "index file `path` to output",
						Destination: &ifile,
					},
				},
				Action: func(cCtx *cli.Context) error {
					arg_len := cCtx.NArg()
					ogg_vorbises := make([]string, arg_len)
					if arg_len > 0 {
						for idx := 0; idx < arg_len; idx++ {
							ogg_vorbises[idx] = cCtx.Args().Get(idx)
						}
					} else {
						return cli.Exit("Argument `ogg-files` required.", 1)
					}

					vorbispak.Merge(out, name, ifile, ogg_vorbises)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
