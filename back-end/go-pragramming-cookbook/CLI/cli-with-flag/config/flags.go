package config

import (
	"flag"
	"fmt"
)

// Config structure for CLI Flag
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

func (c *Config) Setup() {
	// subject 에 대한 긴 버전 및 축약 버전 설정
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")

	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	flag.BoolVar(&c.isAwesome, "ia", false, "is it awesome or what (shorthand)?")

	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")
	flag.IntVar(&c.howAwesome, "ha", 10, "how awesome out of 10?(shorthand)")

	// Custom type에 대한 설정
	flag.Var(&c.countTheWays, "c", "comma seperated list of integers")
}

func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is NOT awesome"
	}

	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s",
		msg, c.howAwesome, c.countTheWays.String(),
	)
	return msg
}
