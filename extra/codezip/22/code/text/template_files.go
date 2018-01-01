package main

import (
	"html/template"
	"log"
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
	t := template.Must(template.ParseGlob("*.tmpl"))
	err := t.ExecuteTemplate(os.Stdout, "todos.tmpl", todoItems)
	if err != nil {
		log.Fatalf("failed executing template: %s", err)
	}
}
