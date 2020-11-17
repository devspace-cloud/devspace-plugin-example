package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

var app = cli.NewApp()

func main() {
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Login",
			Action: func(c *cli.Context) {
				fmt.Println("Logging in...")
				time.Sleep(time.Second * 2)
				fmt.Println("Successfully logged in!")
			},
		},
		{
			Name:  "list",
			Usage: "Lists various things",
			Subcommands: []cli.Command{
				{
					Name:  "env",
					Usage: "Prints environment variables",
					Action: func(c *cli.Context) {
						fmt.Println("Plugin will print the environment variables")
						for _, e := range os.Environ() {
							fmt.Println(e)
						}
					},
				},
			},
		},
		{
			Name:  "print",
			Usage: "Print env vars",
			Subcommands: []cli.Command{
				{
					Name:  "env",
					Usage: "Prints an environment variable",
					Action: func(c *cli.Context) {
						fmt.Print(os.Getenv(c.Args().First()))
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
