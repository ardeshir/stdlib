package main

import (
	"log"
	"strings"
	"unicode"
)

var s = "who,what,when,where,why"

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func dump(i interface{}) {
	log.Printf("%#v", i)
}

func DemoSplit() {
	dump(strings.Split(s, ","))
	dump(strings.SplitN(s, ",", 2))
}

func DemoSplitAfter() {
	dump(strings.SplitAfter(s, ","))
	dump(strings.SplitAfterN(s, ",", 3))
}

func DemoFields() {
	fox := "   The  quick brown Fox jumps    over the lazy Dog."
	dump(strings.Fields(fox))
	dump(strings.FieldsFunc(fox, unicode.IsUpper))
}

func main() {
	DemoSplit()
	DemoSplitAfter()
	DemoFields()
}
