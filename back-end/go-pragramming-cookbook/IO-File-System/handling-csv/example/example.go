package main

import (
	"fmt"
	csv "handling-csv"
	"log"
)

func main() {
	fmt.Println("Read CSV Example: ")
	csv.ReadCSVExample()

	fmt.Println("\nWrite CSV Example")
	csv.WriteCSVBufferExample()

	fmt.Println("\nRead and Write CSV Example")
	fileName := "test.csv"
	books := csv.Books{
		csv.Book{
			Title:       "Go Cookbook",
			Author:      "Aaron Toress",
			Year:        2022,
			Keyword:     "Go",
			Description: "Go Exercise",
		},
		csv.Book{
			Title:       "Go Network Programming",
			Author:      "Adam",
			Year:        2021,
			Keyword:     "golang&network",
			Description: "Network with Go"},
	}

	err := books.WriteFile(fileName)
	if err != nil {
		log.Fatalf("Error for writing files: %s\n", err.Error())
	}

	readingRet, err := csv.ReadCSVByFile(fileName)
	if err != nil {
		log.Fatalf("Error for reading files: %s\n", err.Error())
	}
	fmt.Printf("Reading result: %v\n", readingRet)
}
