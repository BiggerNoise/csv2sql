package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "csv2sql"
	app.Usage = "Converts delimited files to insert statements of a SQL Server dialect"

	app.Action = func(c *cli.Context) error {
		log.Printf("You're a prat")
		return nil
	}

	app.Run(os.Args)
}
