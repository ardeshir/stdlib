package main

import (
	"fmt"
	"log"
	"unicode"
)

var (
	thai = "แล้วฉัน/ผมจะกลับมาใหม่" // I will be right back. http://www.linguanaut.com/english_thai.htm
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoThai() {
	for _, r := range thai {
		fmt.Printf("%c (%U): ", r, r)

		// Query individual runes
		if unicode.IsLetter(r) {
			fmt.Print("IsLetter")
		} else if unicode.IsPunct(r) {
			fmt.Print("IsPunct")
		} else if unicode.IsMark(r) {
			fmt.Print("IsMark")
		}

		// Check if a rune appears in a single range
		fmt.Printf(", Thai?: %t", unicode.Is(unicode.Thai, r))

		// Check if a run appears in multiple ranges (ANY)
		fmt.Printf(", Thai AND Tibetan?: %t", unicode.In(r, unicode.Thai, unicode.Tibetan))

		// Check multiple ranges again (ANY)
		// unicode.In is preferred, because, c'mon, look at that code vs this code.
		// thaiOrTibetan := []*unicode.RangeTable{unicode.Thai, unicode.Tibetan}
		// fmt.Printf(", Thai or Tibetan?: %t", unicode.IsOneOf(thaiOrTibetan, r))

		fmt.Println()
	}
}

func main() {
	DemoThai()
}
