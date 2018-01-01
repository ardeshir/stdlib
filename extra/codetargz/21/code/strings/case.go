package main

import (
	"log"
	"strings"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

var s = "The quick brown Fox jumps over the lazy Dog."

func DemoTitle() {
	log.Println(strings.Title(s))
	log.Println(strings.ToTitle(s))
}

func DemoLower() {
	log.Println(strings.ToLower(s))
}

func DemoUpper() {
	log.Println(strings.ToUpper(s))
}

func main() {
	DemoTitle()
	DemoLower()
	DemoUpper()
}
