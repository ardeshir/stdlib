package main

import (
	"io"
	"log"
	"strings"
)

const (
	example = "The quick brown fox, he likes jumping, you know."
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	lr := io.LimitedReader{strings.NewReader(example), 20}
	buffer := make([]byte, len(example))
	bytesRead, err := lr.Read(buffer)

	// Despite having space, only read 20 bytes
	log.Printf("%s", buffer)
	log.Printf("bytesRead=%d, err=%v", bytesRead, err)

	// Try reading more, won't read anything.
	bytesRead, err = lr.Read(buffer)
	log.Printf("bytesRead=%d, err=%v", bytesRead, err)
}
