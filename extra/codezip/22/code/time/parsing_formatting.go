package main

import (
	"log"
	"time"
)

var (
	layouts = []string{
		time.RFC822,
		time.RFC3339,
		time.Kitchen,
		time.RubyDate,
		"2006-01-_2", // _ to not display leading zeroes
	}
	times = make(chan string, len(layouts))
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoFormat() {
	now := time.Now()
	for _, layout := range layouts {
		formatted := now.Format(layout)
		times <- formatted
		log.Printf("%s + %#v = %#v", now, layout, formatted)
	}
	close(times)
}

func DemoParse() {
	for _, layout := range layouts {
		t := <-times
		parsed, _ := time.Parse(layout, t)
		log.Printf("%#v + %#v = %s", t, layout, parsed)
	}
}

func main() {
	DemoFormat()
	DemoParse()
}
