package main

import (
	"log"
	"os"
	"strings"
)

var s = "All your base are belong to us!"

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	r := strings.NewReader(s)
	log.Println(r.Len())
	r.WriteTo(os.Stdout)
	log.Println(r.Len())
	r.WriteTo(os.Stdout) // It's empty, nothing prints
	r = strings.NewReader(s)

	chunk := make([]byte, 10)
	r.Read(chunk)
	log.Printf("%s", chunk)

	r = strings.NewReader(s)
	// Read a single byte
	b, err := r.ReadByte()
	log.Println(b, err)
	log.Println(r.Len())

	// Nevermind
	r.UnreadByte()
	log.Println(r.Len())
	b, err = r.ReadByte()
	log.Println(b, err)
}
