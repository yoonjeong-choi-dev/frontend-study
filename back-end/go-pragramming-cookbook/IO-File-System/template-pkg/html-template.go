package template_pkg

import (
	"html/template"
	"io"
	"os"
)

type Link struct {
	Url  string
	Name string
}

type HTMLTemplateData struct {
	Title string
	Link  Link
	Items []string
}

const sampleHTMLTemplate = `
<h1>{{.Title}}</h1>

<div>
	List of items
	<ul>
	{{range $index, $item := .Items}}
		<li>{{$item}}</li>
	{{end}}
	</ul>
</div>
<hr/>
<div>
	<a href={{ .Link.Url}}>{{.Link.Name}}</a>
</div>
`

func RunHTMLTemplate(data HTMLTemplateData, writer io.Writer) error {
	t := template.New("html-template-example")

	// parse and validate the template text
	t, err := t.Parse(sampleHTMLTemplate)
	if err != nil {
		return err
	}
	return t.Execute(writer, &data)
}

func SaveHTMLTemplate(fileName string, data HTMLTemplateData) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
	}()

	return RunHTMLTemplate(data, file)
}
