package main

import (
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Polynomial struct {
	U uint32
}

var polynomials = map[string]uint32{
	"ieee":       crc32.IEEE,
	"castagnoli": crc32.Castagnoli,
	"koopman":    crc32.Koopman,
}

func (p *Polynomial) Set(s string) error {
	switch s {
	case "ieee", "castagnoli", "koopman":
		p.U = polynomials[s]
	default:
		var values []string
		for name, _ := range polynomials {
			values = append(values, name)
		}
		return fmt.Errorf("valid values are %s", strings.Join(values, ", "))
	}
	return nil
}

func (p *Polynomial) String() string {
	for name, value := range polynomials {
		if value == p.U {
			return fmt.Sprintf("%s", name)
		}
	}
	panic("not reached")
}

func (p *Polynomial) Table() *crc32.Table {
	return crc32.MakeTable(p.U)
}

var (
	filename   = flag.String("filename", "crc32.go", "The file to checksum")
	streaming  = flag.Bool("streaming", false, "Whether to stream the file instead of reading it all into memory")
	polynomial = &Polynomial{crc32.IEEE}
)

func init() {
	flag.Var(polynomial, "polynomial", "The polynomial to use")
	flag.Parse()
}

func stream(name string) uint32 {
	h := crc32.New(polynomial.Table())
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
	return crc32.Checksum(data, polynomial.Table())
}

func main() {
	var checksum uint32

	if *streaming {
		checksum = stream(*filename)
	} else {
		checksum = simple(*filename)
	}

	log.Printf("the file %s has checksum %#x", *filename, checksum)
}
