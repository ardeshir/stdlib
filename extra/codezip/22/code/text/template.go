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
	t := template.Must(template.New("todos").Parse(`TODO:
{{ range $index, $item := . }}
{{ $index }}: {{ . }}{{ end }}
`))

	t.Execute(os.Stdout, todoItems)
}
