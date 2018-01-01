package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf16"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

type UTF16Reader struct {
	in      io.Reader
	bom     bool
	started bool
}

func NewUTF16Reader(in io.Reader, bom bool) *UTF16Reader {
	return &UTF16Reader{
		in:  in,
		bom: bom,
	}
}

func (r *UTF16Reader) Read(p []byte) (int, error) {
	if !r.started {
		if r.bom {
			// We're assuming little endian, since we used it in the previous example
			bom := make([]byte, 2)
			n, err := r.in.Read(bom)
			if err != nil || n != 2 {
				return n, err
			}
		}
		r.started = true
	}

	// Read some data, deal with the ErrUnexpectedEOF here
	b1 := make([]byte, len(p)/4*4) // We have to read in multiples of 4 bytes
	n, err := io.ReadFull(r.in, b1)
	if err != nil && err != io.ErrUnexpectedEOF {
		return n, err
	}

	// binary.Read some data, make sure it doesn't return ErrUnexpectedEOF, because then it just stops
	b2 := make([]uint16, n/2) // This always rounds down
	err = binary.Read(bytes.NewReader(b1), binary.LittleEndian, b2)
	if err != nil {
		return 0, err
	}

	runes := utf16.Decode(b2)
	bs := []byte(string(runes))
	n = copy(p, bs)
	return n, nil
}

func main() {
	r := NewUTF16Reader(os.Stdin, true)
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalf("failed reading: %s", err)
	}
	os.Stdout.Write(data)
}
