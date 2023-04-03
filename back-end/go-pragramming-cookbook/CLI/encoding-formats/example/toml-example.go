package main

import (
	encoding_formats "encoding-formats"
	"fmt"
)

func TOMLExample() {
	address := encoding_formats.TOMLAddress{
		City:    "Seoul",
		Country: "Korea",
		IsAlone: true,
	}

	data := encoding_formats.TOMLData{
		Name:    "Yoonjeong",
		Age:     31,
		Address: address,
	}

	encoded, err := data.ToTOML()
	if err != nil {
		panic(err)
	}
	fmt.Printf("TOML Marshal=\n%s\n", encoded.String())

	fileName := "toml-example.toml"
	err = data.WriteFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Success to Write. Check the file: %s\n", fileName)

	fileData := encoding_formats.TOMLData{}
	err = fileData.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read file:\n#%v\n", fileData)
}
