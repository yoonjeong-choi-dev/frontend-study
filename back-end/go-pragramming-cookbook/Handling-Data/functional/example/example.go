package main

import (
	"fmt"
	"functional"
)

func printCollection(c functional.Collection) {
	for _, o := range c {
		fmt.Printf("%#v\n", o)
	}
}

func main() {
	c := make(functional.Collection, 0)
	for i := 0; i < 5; i++ {
		o := functional.Item{
			Version: i + 1,
			Data:    fmt.Sprintf("Sample Data %d", i+1),
		}
		c = append(c, o)
	}

	fmt.Println("Initial List:")
	printCollection(c)

	c = functional.Map(c, functional.LowerCaseData)
	fmt.Println("After Lower Case:")
	printCollection(c)

	c = functional.Map(c, functional.IncreaseVersion)
	fmt.Println("After Increase Version")
	printCollection(c)

	c = functional.Filter(c, functional.OldVersionFilter(4))
	fmt.Println("Filtering With version >= 4")
	printCollection(c)
}
