package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.0.0"
const usage = `Usage:
%s [command]
Commands:
	Greet
	Version
`

const greetUsage = `Usage:
%s greet name [flag]
Positional Arguments:
	name
		the name to greet
Flags:
`

type MenuConf struct {
	GoodBye bool
}

// SetupMenu Set basic flag
func (m *MenuConf) SetupMenu() *flag.FlagSet {
	menu := flag.NewFlagSet("menu", flag.ExitOnError)
	menu.Usage = func() {
		fmt.Printf(usage, os.Args[0])
		menu.PrintDefaults()
	}
	return menu
}

// GetSubMenu Set submenu flag
func (m *MenuConf) GetSubMenu() *flag.FlagSet {
	submenu := flag.NewFlagSet("submenu", flag.ExitOnError)
	submenu.BoolVar(&m.GoodBye, "goodbye", false, "Say goodbye instead of hello")
	submenu.Usage = func() {
		fmt.Printf(greetUsage, os.Args[0])
		submenu.PrintDefaults()
	}
	return submenu
}

// Version Function for Main Menu
func (m *MenuConf) Version() {
	fmt.Printf("Version: %s\n", version)
}

// Greet Function for Sub Menu
func (m *MenuConf) Greet(name string) {
	if m.GoodBye {
		fmt.Printf("Goodbye~ %s!\n", name)
	} else {
		fmt.Printf("Hello~ %s!\n", name)
	}
}
