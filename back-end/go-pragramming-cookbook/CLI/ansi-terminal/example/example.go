package main

import (
	ct "ansi-terminal"
	"fmt"
)

func main() {
	r := ct.ColorText{
		TextColor: ct.Red,
		Text:      "This is Red!",
	}

	m := ct.ColorText{
		TextColor: ct.Magenta,
		Text:      "This is Magenta!",
	}

	c := ct.ColorText{
		TextColor: ct.Cyan,
		Text:      "This is Cyan!",
	}

	colors := []ct.ColorText{r, m, c}
	for _, color := range colors {
		fmt.Println(color.String())
	}
}
