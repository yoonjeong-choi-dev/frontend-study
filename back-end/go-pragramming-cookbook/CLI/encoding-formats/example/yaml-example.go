package main

import (
	encoding_formats "encoding-formats"
	"fmt"
)

func YAMLExample() {
	address := encoding_formats.YAMLAddress{
		City:    "Seoul",
		Country: "Korea",
		IsAlone: true,
	}

	data := encoding_formats.YAMLData{
		Name:    "Yoonjeong",
		Age:     31,
		Address: address,
	}

	encoded, err := data.ToYAML()
	if err != nil {
		panic(err)
	}
	fmt.Printf("YAML Marshal=\n%s\n", encoded.String())

	fileName := "yaml-example.yml"
	err = data.WriteFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Success to Write. Check the file: %s\n", fileName)

	fileData := encoding_formats.YAMLData{}
	err = fileData.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read file:\n#%v\n", fileData)
}
