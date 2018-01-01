package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func lines() {
	f, _ := os.Open("scanning.go")
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		log.Printf("line: %s", s.Text())
	}
}

func words() {
	r := strings.NewReader("I just wanna dance with somebody")
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		log.Printf("word: %s", s.Text())
	}
}

func runes() {
	r := strings.NewReader("I just wanna dance with somebody")
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		log.Printf("rune: %s", s.Text())
	}
}

// Basically the `ScanWords` code, altered to split on periods.
func periods(data []byte, atEOF bool) (int, []byte, error) {
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if r != '.' {
			break
		}
	}
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == '.' {
			return i + width, data[start:i], nil
		}
	}
	return 0, nil, nil
}

func custom() {
	f, _ := os.Open("scanning.go")
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(periods)
	for s.Scan() {
		log.Printf("between periods: %s", s.Text())
	}
}

func main() {
	lines()
	words()
	runes()
	custom()
}
