package handling_csv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadCSV(r io.Reader) (Books, error) {
	// csv reader setting
	reader := csv.NewReader(r)
	reader.Comma = Separator
	reader.Comment = Comment

	var output Books

	// ignore header
	_, err := reader.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// Read all body
	for {
		// parse the line
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// validate format
		if len(line) != 5 {
			return nil, errors.New("invalid format")
		}

		year, err := strconv.ParseInt(line[2], 10, 64)
		if err != nil {
			return nil, err
		}

		book := Book{
			Title:       line[0],
			Author:      line[1],
			Year:        int(year),
			Keyword:     line[3],
			Description: line[4],
		}
		output = append(output, book)
	}
	return output, nil
}

func ReadCSVExample() {
	rawString := `
- comment: headers
books title;author;year;keyword;description information
- comment: data parts
Go Cookbook;Aaron Toress;2022;Go;Go Exercise
Go Network Programming;Adam;2021;golang&network;Network with Go
`
	buffer := bytes.NewBufferString(rawString)
	books, err := ReadCSV(buffer)
	if err != nil {
		fmt.Printf("Error for ReadCSV: %s\n", err.Error())
		return
	}

	fmt.Println("Books --------------")
	for _, book := range books {
		fmt.Printf("%#v\n", book)
	}
	return
}

func ReadCSVByFile(fileName string) (books Books, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	books, err = ReadCSV(file)
	return books, err
}
