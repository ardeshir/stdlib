package main

import (
	"io"
	"log"
	"strings"
)

const (
	format = "len(buffer)=%d, min=%d, bytesRead=%d, err=%v, (%s)"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

type Example struct {
	BufferLength int
	MinimumRead  int
	Message      string
}

func ShowExample(ex Example) {
	rd := strings.NewReader(ex.Message)
	buffer := make([]byte, ex.BufferLength)
	bytesRead, err := io.ReadAtLeast(rd, buffer, ex.MinimumRead)
	log.Printf(format, ex.BufferLength, ex.MinimumRead, bytesRead, err, ex.Message)
}

func main() {
	examples := []Example{
		{10, 5, "OK; read less than buf can handle, plenty of data"},
		{100, 75, "Unexpected EOF; buf has space, but ran out of data"},
		{10, 15, "Short buffer; trying to read more than buf can handle"},
	}
	for _, ex := range examples {
		ShowExample(ex)
	}
}
