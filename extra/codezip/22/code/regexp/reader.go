package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var (
	function = regexp.MustCompile(`func (\w+)`)
)

func main() {
	file, err := os.Open("reader.go")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	rr := bufio.NewReader(file)
	log.Println(function.MatchReader(rr))
}
