package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "myapp",
		Usage: "This is a sample CLI application",
	}
	app.Commands = []*cli.Command{
		{
			Name:    "greet",
			Aliases: []string{"g"},
			Usage:   "greet a user",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "lang",
					Value: "english",
					Usage: "language for the greeting",
				},
			},
			Action: func(ctx *cli.Context) error {
				name := "someone"
				if ctx.NArg() > 0 {
					name = ctx.Args().Get(0)
				}
				lang := ctx.String("lang")
				switch lang {
				case "spanish":
					fmt.Println("Hola", name)
				case "chinese":
					fmt.Println("Ni Hao, name")
				default:
					fmt.Println("Hello", name)
				}
				return nil
			},
		},
		{
			Name:    "query",
			Aliases: []string{"q"},
			Usage:   "Query a user",
			Action: func(ctx *cli.Context) error {
				RunQuery()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
