package ansi_terminal

import "fmt"

type Color int

// Define Color for Text
const (
	ColorNone = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Black Color = -1
)

type ColorText struct {
	TextColor Color
	Text      string
}

func (t *ColorText) String() string {
	if t.TextColor == ColorNone {
		return t.Text
	}

	// value: 터미널에서 색을 표현하기 위한 정수값
	value := 30
	if t.TextColor != Black {
		value += int(t.TextColor)
	}
	return fmt.Sprintf("\033[0;%dm%s\033[0m", value, t.Text)
}
