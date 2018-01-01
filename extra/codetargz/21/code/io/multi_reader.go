package main

import (
	"io"
	"log"
	"strings"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	// Make our inputs
	a := strings.NewReader(strings.Repeat("A", 5))
	b := strings.NewReader(strings.Repeat("B", 5))
	c := strings.NewReader(strings.Repeat("C", 5))

	// Read ALL THE THINGS
	mr := io.MultiReader(a, b, c)

	// Read A
	buffer := make([]byte, 20)
	n1, err := mr.Read(buffer)
	log.Printf("%v", buffer)
	log.Printf("n1=%d, err=%v", n1, err)

	// Read B
	n2, err := mr.Read(buffer[n1:])
	log.Printf("%v", buffer)
	log.Printf("n2=%d, err=%v", n2, err)

	// Read C
	n3, err := mr.Read(buffer[(n1 + n2):])
	log.Printf("%v", buffer)
	log.Printf("n3=%d, err=%v", n3, err)

	// EOF
	n4, err := mr.Read(buffer[(n1 + n2 + n3):])
	log.Printf("%v", buffer)
	log.Printf("n4=%d, err=%v", n4, err)
}
