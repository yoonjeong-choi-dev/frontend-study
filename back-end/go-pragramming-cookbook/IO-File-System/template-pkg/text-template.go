package template_pkg

import (
	"io"
	"strings"
	"text/template"
)

type TextTemplateData struct {
	Condition      bool
	StringVariable string
	Items          []string
	Words          string
	OtherVariable  string
}

const sampleTextTemplate = `
	This template demonstrates printing a {{ .StringVariable | printf "%#v" }}.

	[Conditional Statement]
	{{if .Condition}}
    If Condition is set, this statement is printed
    {{else}}
    Otherwise, this statement is printed instead
    {{end}}
	
	[Loop Statement]
	Let's iterate over an array of strings:
	{{range $index, $item := .Items}}
		{{$index}} : {{$item}}
	{{end}}

	[Import Function]
	Let's import "strings.Split" function and iterate the result of the function
    {{ range $index, $item := yjSplit .Words ","}}
        {{$index}} : {{$item}}
    {{end}}

	[Embedding another template]
	We can also embed another template
	{{ block "block_example" .}}
		No Embedding Template whose name is 'block_example'
	{{end}}
`

const embeddedTemplate = `
	{{define "block_example"}}
	{{.OtherVariable}}
	{{end}}
`

func RunTextTemplate(data TextTemplateData, writer io.Writer) error {
	functionMapToImport := template.FuncMap{
		"yjSplit": strings.Split,
	}

	t := template.New("simple-template-example")
	t = t.Funcs(functionMapToImport)

	// parse and validate the template text
	t, err := t.Parse(sampleTextTemplate)
	if err != nil {
		return err
	}

	// parse and validate the block template
	t, err = t.Parse(embeddedTemplate)
	if err != nil {
		return err
	}

	return t.Execute(writer, &data)
}
