package main

import (
	"html"
	"log"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

func main() {
	raw := []string{
		"hello",
		"<i>hello</i>",
		"alert('hello');",
		"foo & bar",
		`"how are you?" he asked.`,
	}

	log.Println("html.EscapeString")
	for _, s := range raw {
		log.Printf("\t%s -> %s", s, html.EscapeString(s))
	}

	log.Println("html.UnescapeString(html.EscapeString)")
	for _, s := range raw {
		flipped := html.UnescapeString(html.EscapeString(s))
		log.Printf("\t%s -> %s", s, flipped)
	}

	escaped := []string{
		"&#225;",
		"&raquo;",
		"&middot;",
		"&lt;i&gt;hello&lt;/i&gt;",
	}

	log.Println("html.UnescapeString")
	for _, s := range escaped {
		log.Printf("\t%s -> %s", s, html.UnescapeString(s))
	}
}
