package main

import (
	"log"
	"strings"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

var s = "  \n all the spaces \t  "

func dump(i interface{}) {
	log.Printf("%#v", i)
}

func DemoTrim() {
	cutset := " \t\n"
	dump(strings.Trim(s, cutset))
	dump(strings.TrimLeft(s, cutset))
	dump(strings.TrimRight(s, cutset))
	dump(strings.TrimSpace(s))
}

func DemoPrefixSuffix() {
	s2 := "The Go Programming Language"
	dump(s2)
	s2 = strings.TrimPrefix(s2, "The Go ")
	dump(s2)
	s2 = strings.TrimSuffix(s2, " Language")
	dump(s2)
}

func onlySpaces(r rune) bool {
	return r == ' '
}

func DemoTrimFunc() {
	dump(strings.TrimFunc(s, onlySpaces))
	dump(strings.TrimLeftFunc(s, onlySpaces))
	dump(strings.TrimRightFunc(s, onlySpaces))
}

func main() {
	dump(s)
	DemoTrim()
	DemoPrefixSuffix()
	DemoTrimFunc()
}
