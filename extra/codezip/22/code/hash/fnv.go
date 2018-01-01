package main

import (
	"flag"
	"hash/fnv"
	"io"
	"log"
	"os"
)

var (
	filename = flag.String("filename", "fnv.go", "The file to checksum")
	_64bit   = flag.Bool("64", false, "Use the 64-bit interface")
)

func runHash(name string, w io.Writer) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("failed opening %s: %s", name, err)
	}
	defer file.Close()
	io.Copy(w, file)
}

func hash64(name string) uint64 {
	h := fnv.New64()
	runHash(name, h)
	return h.Sum64()
}

func hash32(name string) uint32 {
	h := fnv.New32()
	runHash(name, h)
	return h.Sum32()
}

func main() {
	flag.Parse()
	if *_64bit {
		h := hash64(*filename)
		log.Printf("the file %s has hash %#x", *filename, h)
	} else {
		h := hash32(*filename)
		log.Printf("the file %s has hash %#x", *filename, h)
	}
}
