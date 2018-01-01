package main

import (
	"io"
	"log"
	"runtime"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	runtime.GOMAXPROCS(8)
}

func Write(wr io.WriteCloser) {
	lyrics := []string{
		"I come home in the morning light",
		"My mother says when you gonna live your life right",
		"Oh mother dear we're not the fortunate ones",
		"And girls they want to have fun",
		"Oh girls just want to have fun",
	}

	for _, line := range lyrics {
		io.WriteString(wr, line)
	}
	wr.Close() // We're done, signal EOF
}

func main() {
	rd, wr := io.Pipe()
	go Write(wr)
	for {
		buf := make([]byte, 32)
		n, err := rd.Read(buf)
		log.Printf("buf: %s", buf)
		log.Printf("n=%d, err=%v", n, err)
		if err == io.EOF {
			break
		}
	}
}
