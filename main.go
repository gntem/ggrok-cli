package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vjeantet/grok"
)

func main() {
	var str string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "parse",
				Aliases:     []string{"p"},
				Usage:       "Parse input str",
				Destination: &str,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	g, _ := grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})

	values, _ := g.Parse("%{COMMONAPACHELOG}", str)

	for k, v := range values {
		fmt.Printf("%+15s: %s\n", k, v)
	}
}
