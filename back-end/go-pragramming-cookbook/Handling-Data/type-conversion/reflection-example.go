package main

import "fmt"

func CheckType(o interface{}) {
	var t string
	switch o.(type) {
	case string:
		t = "string"
	case int:
		t = "int"
	default:
		t = "UNKNOWN"
	}

	fmt.Printf("%v - type: %s\n", o, t)
}

func InterfaceConversion() {
	CheckType("test string")
	CheckType(1234)
	CheckType(true)

	var i interface{}
	i = "stringValue"
	CheckType(i)

	if val, ok := i.(string); ok {
		fmt.Printf("%v has string value: %s\n", i, val)
	}

	if _, ok := i.(int); !ok {
		fmt.Printf("%v cannot be converted to int type", i)
	}
}
