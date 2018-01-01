package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func buffer() *bytes.Buffer {
	var buf bytes.Buffer
	buf.WriteString("I'm writing ")
	buf.WriteString("strings ")
	buf.WriteString("to this buffer ")
	buf.WriteString("and we'll copy it to os.Stdout.\n")
	return &buf
}

func DemoCopy() {
	buf := buffer()
	log.Printf("copying %d bytes to os.Stdout", buf.Len())
	io.Copy(os.Stdout, buf)
}

func DemoCopyN() {
	buf := buffer()
	n := int64(32)
	log.Printf("have %d bytes, only copying %d to os.Stdout", buf.Len(), n)
	io.CopyN(os.Stdout, buf, n)
	os.Stdout.Write([]byte{'\n'})
}

func BufferFun() {
	buf := buffer()
	n, _ := io.CopyN(os.Stdout, buf, 32)
	nn, _ := io.Copy(os.Stdout, buf)
	log.Printf("copied %d and then %d bytes to os.Stdout", n, nn)
}

func main() {
	DemoCopy()
	DemoCopyN()
	BufferFun()
}
