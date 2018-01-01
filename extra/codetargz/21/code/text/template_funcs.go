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
	tmpl.Funcs(map[string]interface{}{
		"inc": func(a, b int) int {
			return a + b
		},
	})
	t := template.Must(tmpl.Parse(`TODO:
{{ range $index, $item := . }}
{{ inc $index 1 }}: {{ $item }}{{ end }}
`))

	t.Execute(os.Stdout, todoItems)
}
