package main

import (
	"log"
	"regexp"
)

var (
	eqn = "3x * 2y - 9 / 3 / 5 * 5"
	mul = regexp.MustCompile(`\w+ \* \w+`)
	div = regexp.MustCompile(`\w+ / \w+`)
)

func main() {
	fmul := mul.FindStringIndex(eqn)
	log.Println(fmul, eqn[fmul[0]:fmul[1]])

	divs := div.FindAllStringIndex(eqn, -1)
	log.Println("divs", divs)
	for index, pair := range divs {
		log.Printf("match %d: %s", index, eqn[pair[0]:pair[1]])
	}
}
