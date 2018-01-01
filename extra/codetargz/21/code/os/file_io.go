package main

import (
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoCreate() {
	f, err := os.Create("demo.txt") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer f.Close() // Make sure to close the file when you're done

	n, err := f.WriteString(`"And live from New York, it's Saturday Night!" - Cast of SNL`)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	log.Printf("wrote %d bytes to %s", n, f.Name())
}

func DemoOpenFile() {
	// OpenFile lets you customize whether the file is truncated, must exist, or must not exist, etc
	// Open is your basic way to open a file for reading, but we need to write.
	f, err := os.OpenFile("demo.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer f.Close()

	n, err := f.WriteString("\nSince 1985\n")
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	log.Printf("wrote another %d bytes to %s", n, f.Name())
}

func DemoWriteAt() {
	// In DemoOpenFile, we wrote the wrong date, let's fix that
	f, err := os.OpenFile("demo.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer f.Close()

	n, err := f.WriteAt([]byte{'7'}, 69)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	log.Printf("wrote another %d bytes to %s", n, f.Name())
}

func DemoRead() {
	f, err := os.Open("demo.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("failed reading %s: %s", f.Name(), err)
	}
	log.Printf("contents:\n%s", data)
}

func main() {
	DemoCreate()
	DemoRead()
	DemoOpenFile()
	DemoRead()
	DemoWriteAt()
	DemoRead()
}
