package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	c := MenuConf{}
	menu := c.SetupMenu()
	if err := menu.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing params %s, error: %v", os.Args[1:], err)
		os.Exit(1)
	}

	argsLen := len(os.Args)

	if argsLen > 1 {
		// Process sub command
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			submenu := c.GetSubMenu()
			if argsLen < 3 {
				submenu.Usage()
				os.Exit(1)
			}
			if argsLen > 3 {
				if err := submenu.Parse(os.Args[3:]); err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing parmas %s, error: %v",
						os.Args[3:], err,
					)
					os.Exit(1)
				}
			}

			c.Greet(os.Args[2])
		default:
			fmt.Fprintln(os.Stderr, "Invalid command")
			menu.Usage()
			os.Exit(1)
		}
	} else {
		menu.Usage()
		os.Exit(1)
	}
}
