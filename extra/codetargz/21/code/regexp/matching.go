package main

import (
	"bytes"
	"log"
	"regexp"
)

var (
	universes = regexp.MustCompile(`(batman and robin)|(thor and loki)`)
	heroes    = "batman and robin"
)

func main() {
	log.Println(universes.MatchString(heroes))
	log.Println(universes.Match([]byte(heroes)))
	rr := bytes.NewBufferString(heroes)
	log.Println(universes.MatchReader(rr))

	log.Println(universes.MatchString("batman and loki"))
}
