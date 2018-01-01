package main

import (
	"io"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	io.WriteString(w, "Hello, twice!!\n")
}
