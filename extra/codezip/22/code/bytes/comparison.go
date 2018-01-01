package main

import (
	"bytes"
	"log"
)

func DemoCompare(a, b []byte) {
	if c := bytes.Compare(a, b); c == -1 {
		log.Printf("%s is less than %s", a, b)
	} else if c == 1 {
		log.Printf("%s is greater than %s", a, b)
	} else {
		log.Printf("%s and %s are equal", a, b)
	}
}

func DemoEqual(a, b []byte) {
	if bytes.Equal(a, b) {
		log.Printf("%s and %s are equal", a, b)
	} else {
		log.Printf("%s and %s are NOT equal", a, b)
	}
}

func DemoEqualFold(a, b []byte) {
	if bytes.EqualFold(a, b) {
		log.Printf("%s and %s are equal", a, b)
	} else {
		log.Printf("%s and %s are NOT equal", a, b)
	}
}

func main() {
	golang := []byte("golang")
	gOlaNg := []byte("gOlaNg")
	haskell := []byte("haskell")

	DemoCompare(golang, golang)
	DemoCompare(golang, haskell)
	DemoCompare(haskell, golang)

	DemoEqual(golang, golang)
	DemoEqual(golang, haskell)

	DemoEqualFold(golang, gOlaNg)
	DemoEqualFold(golang, golang)
}
