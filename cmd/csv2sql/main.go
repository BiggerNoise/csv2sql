package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/biggernoise/csv2sql"
	"github.com/codegangsta/cli"
)

type Config struct {
	csv2sql.MsSqlWriterConfig
	Delimeter string
}

func main() {
	app := cli.NewApp()
	app.Name = "csv2sql"
	app.Usage = "Converts delimited files to insert statements of a SQL Server dialect"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "input", Usage: "Path to the input source"},
		cli.StringFlag{Name: "config", Usage: "Path to the config File"},
	}
	app.Action = action

	app.Run(os.Args)
}

func action(c *cli.Context) {

	cf, err := os.Open(c.String("config"))
	if err != nil {
		log.Printf("unable to open config file")
		os.Exit(1)
	}

	dec := json.NewDecoder(cf)
	var config Config
	err = dec.Decode(&config)
	if err != nil {
		log.Printf("unable to parse config file: %s", err)
		os.Exit(1)
	}

	input, err := os.Open(c.String("input"))
	if err != nil {
		log.Printf("unable to open input file")
		os.Exit(1)
	}

	var delimeter rune
	for _, c := range config.Delimeter {
		delimeter = c
		break
	}
	reader := csv2sql.NewCsvReader(&csv2sql.CsvReaderConfig{Delimeter: delimeter}, input)
	writer := csv2sql.NewMsSqlWriter(&(config.MsSqlWriterConfig), os.Stdout)

	for record, err := reader.Read(); err == nil; record, err = reader.Read() {
		writer.Write(record)
	}
	writer.Done()

	return
}
