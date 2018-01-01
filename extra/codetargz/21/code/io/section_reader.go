package main

import (
	"bytes"
	"io"
	"log"
)

var (
	s = "The quick brown fox, he likes jumping, you know."
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	// Build the block of data
	data := make([]byte, 0, 30)
	data = append(data, bytes.Repeat([]byte{'A'}, 10)...)
	data = append(data, bytes.Repeat([]byte{'B'}, 10)...)
	data = append(data, bytes.Repeat([]byte{'C'}, 10)...)

	// Create some SectionReaders to read the 3 sections
	r := bytes.NewReader(data)
	ar := io.NewSectionReader(r, 0, 10)
	br := io.NewSectionReader(r, 10, 10)
	cr := io.NewSectionReader(r, 20, 10)

	buf := make([]byte, 10)

	// Read the A section
	n, err := ar.Read(buf)
	log.Printf("buf: %s", buf)
	log.Printf("n=%d, err=%v", n, err)

	// Read the B section
	n, err = br.Read(buf)
	log.Printf("buf: %s", buf)
	log.Printf("n=%d, err=%v", n, err)

	// Read the C section
	n, err = cr.Read(buf)
	log.Printf("buf: %s", buf)
	log.Printf("n=%d, err=%v", n, err)
}
