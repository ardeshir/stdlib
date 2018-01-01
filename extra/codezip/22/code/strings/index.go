package main

import (
	"log"
	"strings"
	"unicode"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

var s = "Go, The Standard Library"

// Find specific things
func DemoIndex() {
	needles := []string{",", "t", "The", "x"}
	for _, needle := range needles {
		index := strings.Index(s, needle)
		log.Printf("Index(%#v) %d", needle, index)
	}
}

// Search for any unicode code points
func DemoIndexAny() {
	needles := []string{",thx", "ray"}
	for _, needle := range needles {
		index := strings.IndexAny(s, needle)
		log.Printf("IndexAny(%#v) %d", needle, index)
	}
}

// Search for a specific byte
func DemoIndexByte() {
	needles := []byte{',', 'y'}
	for _, needle := range needles {
		index := strings.IndexByte(s, needle)
		log.Printf("IndexByte(%q) %d", needle, index)
	}
}

func nonAlphaNumeric(r rune) bool {
	switch {
	case 48 <= r && r <= 57: // numbers
		return false
	case 97 <= r && r <= 122: // lowercase
		return false
	case 65 <= r && r <= 90: // uppercase
		return false
	}
	return true
}

// Use a function
func DemoIndexFunc() {
	funcs := []struct {
		name string
		f    func(rune) bool
	}{
		{"nonAlphaNumeric", nonAlphaNumeric},
		{"unicode.IsLower", unicode.IsDigit},
		{"unicode.IsLower", unicode.IsLower},
	}
	for _, f := range funcs {
		index := strings.IndexFunc(s, f.f)
		log.Printf("IndexFunc(%#v) %d", f.name, index)
	}
}

// Find a specific rune
func DemoIndexRune() {
	runes := []rune{'a', ' ', '.'}
	for _, r := range runes {
		index := strings.IndexRune(s, r)
		log.Printf("IndexRune(%q) %d", r, index)
	}
}

// Find the last index of a substring
func DemoLastIndex() {
	needles := []string{"a", "r", "y", "\t"}
	for _, needle := range needles {
		index := strings.LastIndex(s, needle)
		log.Printf("LastIndex(%#v) %d", needle, index)
	}
}

// Find the last index of any of the given unicode code points
func DemoLastIndexAny() {
	needles := []string{",thx", "ray"}
	for _, needle := range needles {
		index := strings.LastIndexAny(s, needle)
		log.Printf("LastIndexAny(%#v) %d", needle, index)
	}
}

// Use a func to find the last index of something
func DemoLastIndexFunc() {
	funcs := []struct {
		name string
		f    func(rune) bool
	}{
		{"nonAlphaNumeric", nonAlphaNumeric},
		{"unicode.IsLower", unicode.IsDigit},
		{"unicode.IsLower", unicode.IsLower},
	}
	for _, f := range funcs {
		index := strings.LastIndexFunc(s, f.f)
		log.Printf("LastIndexFunc(%#v) %d", f.name, index)
	}
}

func main() {
	log.Printf("haystack: %#v", s)

	DemoIndex()
	DemoIndexAny()
	DemoIndexByte()
	DemoIndexFunc()
	DemoIndexRune()
	DemoLastIndex()
	DemoLastIndexAny()
	DemoLastIndexFunc()
}
