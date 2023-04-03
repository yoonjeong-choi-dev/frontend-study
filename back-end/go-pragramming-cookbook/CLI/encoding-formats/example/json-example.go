package main

import (
	encoding_formats "encoding-formats"
	"fmt"
)

func JSONExample() {
	address := encoding_formats.JSONAddress{
		City:    "Seoul",
		Country: "Korea",
		IsAlone: true,
	}

	data := encoding_formats.JSONData{
		Name:    "Yoonjeong",
		Age:     31,
		Address: address,
	}

	encoded, err := data.ToJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON Marshal=\n%s\n", encoded.String())

	fileName := "json-example.json"
	err = data.WriteFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Success to Write. Check the file: %s\n", fileName)

	fileData := encoding_formats.JSONData{}
	err = fileData.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read file:\n#%v\n", fileData)
}
