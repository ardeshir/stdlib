package main

import (
	"log"
	"strings"
)

var s = "Go, The Standard Library"

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

// Look for exact matches
func DemoContains() {
	needles := []string{"Library", "standard", "Standard"}
	for _, needle := range needles {
		found := strings.Contains(s, needle)
		log.Printf("Contains(%#v) %t", needle, found)
	}
}

// Look for any of the unicode code points
func DemoContainsAny() {
	sets := []string{"aeiou", "zyx", "\t\r"}
	for _, set := range sets {
		found := strings.ContainsAny(s, set)
		log.Printf("ContainsAny(%#v) %t", set, found)
	}
}

func DemoContainsRune() {
	runes := []rune{'a', ' ', '.'}
	for _, rune := range runes {
		found := strings.ContainsRune(s, rune)
		log.Printf("ContainsRune(%q) %t", rune, found)
	}
}

// Count substrings
func DemoCount() {
	needles := []string{"", "a", ", "}
	for _, needle := range needles {
		count := strings.Count(s, needle)
		log.Printf("Count(%#v) %d", needle, count)
	}
}

// Is it equal ignoring unicode case
func DemoEqualFold() {
	ts := []string{s, strings.ToUpper(s), strings.ToLower(s)}
	for _, t := range ts {
		equal := strings.EqualFold(s, t)
		log.Printf("EqualFold(%#v) %t", t, equal)
	}
}

// Check for prefixes
func DemoHasPrefix() {
	prefixes := []string{"Go", "GO", "Go, "}
	for _, prefix := range prefixes {
		has := strings.HasPrefix(s, prefix)
		log.Printf("HasPrefix(%#v) %t", prefix, has)
	}
}

// Check for suffixes
func DemoHasSuffix() {
	suffixes := []string{"Library", "", "Standard"}
	for _, suffix := range suffixes {
		has := strings.HasSuffix(s, suffix)
		log.Printf("HasSuffix(%#v) %t", suffix, has)
	}
}

func main() {
	log.Printf("haystack: %#v", s)

	DemoContains()
	DemoContainsAny()
	DemoContainsRune()
	DemoCount()
	DemoEqualFold()
	DemoHasPrefix()
	DemoHasSuffix()
}
