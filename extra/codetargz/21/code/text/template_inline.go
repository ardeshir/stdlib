package main

import (
	"html/template"
	"os"
)

var (
	todoItems = []string{
		"cut the grass",
		"pick up milk",
		"feed the dog",
	}
)

func main() {
	tmpl := template.New("todos")
	t := template.Must(tmpl.Parse(`{{ define "todo" }}- {{ . }}{{ end }}TODO:
{{ range $index, $item := . }}
{{ template "todo" $item }}{{ end }}
`))

	t.Execute(os.Stdout, todoItems)
}
