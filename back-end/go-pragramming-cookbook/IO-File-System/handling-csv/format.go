package handling_csv

type Book struct {
	Title       string
	Author      string
	Year        int
	Keyword     string
	Description string
}

type Books []Book

var (
	Separator = ';'
	Comment   = '-'
)
