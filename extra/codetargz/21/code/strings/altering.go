package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

var s = "red green blue"

func DemoJoin() {
	fields := strings.Fields(s)
	log.Println(strings.Join(fields, ","))
	log.Println(strings.Join(fields, ":"))
	log.Println(strings.Join(fields, ""))
}

func rot13(r rune) rune {
	switch {
	case 65 <= r && r <= 90:
		return 65 + ((r-65)+13)%26
	case 97 <= r && r <= 122:
		return 97 + ((r-97)+13)%26
	default:
		return r
	}
}

func DemoMap() {
	log.Println(strings.Map(unicode.ToUpper, s))
	mapped := strings.Map(func(r rune) rune {
		switch r {
		case 'e':
			return -1
		default:
			return r + 1
		}
	}, s)
	log.Println(mapped)
	log.Println(strings.Map(rot13, s))
}

func DemoRepeat() {
	log.Println(strings.Repeat("-", len(s)))
}

func DemoReplace() {
	log.Println(strings.Replace(s, "e", "!", 1))
	log.Println(strings.Replace(s, "e", "!", -1))
}

func DemoReplacer() {
	r := strings.NewReplacer("e", "E")
	log.Println(r.Replace(s))
	r.WriteString(os.Stdout, s)
}

func main() {
	DemoJoin()
	DemoMap()
	DemoRepeat()
	DemoReplace()
	DemoReplacer()
}
