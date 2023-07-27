package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type appConfig struct {
	format   string
	dataFile string
}

func parseArgs() (appConfig, error) {
	config := appConfig{}

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			`Usage: %s [flags] [add chore, ...|complete #]
    add         add comma-separated chores
    complete    complete designated chore
    list        list all chores

Flags:
`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.StringVar(&config.format, "format", "json", "serialize format(json, gob, protobuf)")
	flag.StringVar(&config.dataFile, "file", "housework.db", "data file")

	flag.Parse()

	return config, nil
}

func main() {
	config, err := parseArgs()
	if err != nil {
		fmt.Printf("failed to parse: %s\n", err.Error())
		os.Exit(1)
	}

	app, err := NewHouseworkApp(config.format, config.dataFile)
	if err != nil {
		fmt.Printf("failed to init app: %s\n", err.Error())
		os.Exit(1)
	}

	switch strings.ToLower(flag.Arg(0)) {
	case "list":
		err = app.printList()
	case "add":
		err = app.add(strings.Join(flag.Args()[1:], " "))
	case "complete":
		err = app.complete(flag.Arg(1))
	default:
		err = errors.New("unsupported command")
	}

	if err != nil {
		fmt.Printf("failed to process: %s\n", err.Error())
		os.Exit(1)
	}

	if flag.Arg(0) != "list" {
		err = app.printList()
		if err != nil {
			fmt.Printf("failed to list: %s\n", err.Error())
			os.Exit(1)
		}
	}
}
