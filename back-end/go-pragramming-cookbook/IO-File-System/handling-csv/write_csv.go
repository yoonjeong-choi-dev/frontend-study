package handling_csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func (books *Books) ToCSV(w io.Writer) error {
	writer := csv.NewWriter(w)
	writer.Comma = Separator

	// Writer Header
	err := writer.Write([]string{"Title", "Author", "Year", "Keyword", "Description"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err := writer.Write([]string{
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			book.Keyword,
			book.Description,
		})

		if err != nil {
			return err
		}
	}

	writer.Flush()
	return writer.Error()
}

func (books *Books) WriteCSVBuffer() (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}
	err := books.ToCSV(buffer)
	return buffer, err
}

func (books *Books) WriteFile(fileName string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	err = books.ToCSV(file)
	if err != nil {
		return err
	}

	return nil
}

func WriteCSVBufferExample() {
	books := Books{
		Book{
			Title:       "Go Cookbook",
			Author:      "Aaron Toress",
			Year:        2022,
			Keyword:     "Go",
			Description: "Go Exercise",
		},
		Book{
			Title:       "Go Network Programming",
			Author:      "Adam",
			Year:        2021,
			Keyword:     "golang&network",
			Description: "Network with Go"},
	}

	buffer, err := books.WriteCSVBuffer()
	if err != nil {
		fmt.Printf("Error for WriteCSV: %s\n", err.Error())
		return
	}

	fmt.Printf("Buffer: %s\n", buffer.String())
}
