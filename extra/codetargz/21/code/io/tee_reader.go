package main

import (
	"bytes"
	"io"
	"log"
	"strings"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	s := strings.NewReader("Get to the choppa!")
	var buf bytes.Buffer
	tr := io.TeeReader(s, &buf)
	b := make([]byte, s.Len())
	n, err := tr.Read(b)
	log.Printf("buf: %s", &buf)
	log.Printf("  b: %s", b)
	log.Printf("n=%d, err=%v", n, err)
}
