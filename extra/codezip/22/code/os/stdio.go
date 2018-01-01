package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoStdin() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed reading stdin: %s", err)
	}
	log.Printf("read %d from stdin: %s", len(input), input)
}

func DemoDevNull() {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		log.Fatalf("failed opening null device: %s", err)
	}
	defer devNull.Close()
	io.WriteString(devNull, "This is going nowhere\n")
}

func main() {
	io.WriteString(os.Stdout, "This is stdout\n")
	io.WriteString(os.Stderr, "This is stderr\n")
	DemoDevNull()
	DemoStdin()
}
