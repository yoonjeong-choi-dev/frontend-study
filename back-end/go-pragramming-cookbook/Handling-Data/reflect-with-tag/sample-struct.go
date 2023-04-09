package reflect_with_tag

import "fmt"

type SampleStruct struct {
	Name  string `serialize:"name"`
	City  string `serialize:"city"`
	State string
	Misc  string `serialize:"-"`
	Year  int    `serialize:"year"`
}

func encodeAndDecodeExample(s SampleStruct) error {
	res, err := SerializeStringsStruct(&s)
	if err != nil {
		return nil
	}

	fmt.Printf("Original struct: %#v\n", s)
	fmt.Printf("Serialze Reuslt: %s\n", res)

	decoded := SampleStruct{}
	if err := DeserializeStringsStruct(res, &decoded); err != nil {
		return err
	}

	fmt.Printf("Deserialize Result: %#v\n", decoded)
	return nil
}

func EmptyStructExample() error {
	fmt.Println("Empty Struct Example")
	s := SampleStruct{}

	return encodeAndDecodeExample(s)
}

func FullStructExample() error {
	fmt.Println("Full Struct Example")

	s := SampleStruct{
		Name:  "Yoonjeong",
		City:  "Seoul",
		State: "Gangnam",
		Misc:  "miscellaneous",
		Year:  1993,
	}

	return encodeAndDecodeExample(s)
}
