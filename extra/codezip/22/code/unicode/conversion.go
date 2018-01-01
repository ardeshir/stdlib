package main

import (
	"flag"
	"fmt"
	"log"
	"unicode"
)

var (
	toggle string
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")

	flag.StringVar(&toggle, "toggle", "MýrDalSjökuLL", "toggle the case of each unicode rune")
	flag.Parse()
}

func main() {
	toggled := make([]rune, len(toggle))
	for index, r := range toggle {
		if unicode.IsUpper(r) {
			toggled[index] = unicode.ToLower(r)
		} else {
			toggled[index] = unicode.ToUpper(r)
		}
	}
	fmt.Printf("original: %s\ntoggled: %s\n", toggle, string(toggled))
}
