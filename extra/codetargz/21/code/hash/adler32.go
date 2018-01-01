package main

import (
	"flag"
	"hash/adler32"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	filename  = flag.String("filename", "adler32.go", "The file to checksum")
	streaming = flag.Bool("streaming", false, "Whether to stream the file instead of reading it all into memory")
)

func stream(name string) uint32 {
	h := adler32.New()
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("failed opening %s: %s", name, err)
	}
	defer file.Close()
	io.Copy(h, file)
	return h.Sum32()
}

func simple(name string) uint32 {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalf("failed reading %s: %s", name, err)
	}
	return adler32.Checksum(data)
}

func main() {
	flag.Parse()
	var checksum uint32

	if *streaming {
		checksum = stream(*filename)
	} else {
		checksum = simple(*filename)
	}

	log.Printf("the file %s has checksum %#x", *filename, checksum)
}
