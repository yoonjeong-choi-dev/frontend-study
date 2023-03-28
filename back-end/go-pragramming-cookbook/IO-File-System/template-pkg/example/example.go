package main

import (
	"fmt"
	"os"
	template_pkg "template-pkg"
)

func main() {
	fmt.Println("Text based Template Example")
	data := template_pkg.TextTemplateData{
		Condition:      true,
		StringVariable: "This is Title",
		Items:          []string{"item1", "item2", "item3"},
		Words:          "go,java,cpp,python,javascript",
		OtherVariable:  "This is the content for embedded templates",
	}
	if err := template_pkg.RunTextTemplate(data, os.Stdout); err != nil {
		fmt.Printf("Error for RunTextTemplate: %s\n", err.Error())
	}

	fmt.Println("\nHTML based Template Example")
	htmlData := template_pkg.HTMLTemplateData{
		Title: "HTML Template Test",
		Link: template_pkg.Link{
			Url:  "https://www.google.com",
			Name: "Google",
		},
		Items: []string{"Go", "Javascript", "Python"},
	}
	if err := template_pkg.RunHTMLTemplate(htmlData, os.Stdout); err != nil {
		fmt.Printf("Error for RunHTMLTemplate: %s\n", err.Error())
	}

	fmt.Println("\nHTML based Template Example - Save file")
	if err := template_pkg.SaveHTMLTemplate("test.html", htmlData); err != nil {
		fmt.Printf("Error for RunHTMLTemplate: %s\n", err.Error())
	}

}
