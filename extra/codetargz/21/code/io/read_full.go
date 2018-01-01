package main

import (
	"io"
	"log"
	"strings"
)

const (
	format = "len(buffer)=%d, bytesRead=%d, err=%v, (%s)"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

type Example struct {
	BufferLength int
	Message      string
}

func ShowExample(ex Example) {
	rd := strings.NewReader(ex.Message)
	buffer := make([]byte, ex.BufferLength)
	bytesRead, err := io.ReadFull(rd, buffer)
	log.Printf("%v", buffer)
	log.Printf(format, ex.BufferLength, bytesRead, err, ex.Message)
}

func main() {
	examples := []Example{
		{10, "OK; filled up buf, plenty of data"},
		{55, "Unexpected EOF; buf has space, but ran out of data"},
	}
	for _, ex := range examples {
		ShowExample(ex)
	}
}
