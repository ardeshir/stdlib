package main

import (
	"bufio"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func openFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("failed opening %s for writing: %s", name, err)
	}
	return file
}

func doWriteByte(w *bufio.Writer) {
	if err := w.WriteByte('G'); err != nil {
		log.Fatalf("failed writing a byte: %s", err)
	}
}

func doWriteRune(w *bufio.Writer) {
	if written, err := w.WriteRune(rune('o')); err != nil {
		log.Fatalf("failed writing a rune: %s", err)
	} else {
		log.Printf("Wrote rune in %d bytes", written)
	}
}

func doWriteString(w *bufio.Writer) {
	written, err := w.WriteString(", The Standard Library\n")
	if err != nil {
		log.Fatalf("failed writing string: %s", err)
	}
	log.Printf("Wrote string in %d bytes", written)
}

func main() {
	file := openFile("bufio.out")
	defer file.Close()

	bw := bufio.NewWriter(file)

	// Remember to Flush!
	defer bw.Flush()

	doWriteByte(bw)
	doWriteRune(bw)
	doWriteString(bw)
}
