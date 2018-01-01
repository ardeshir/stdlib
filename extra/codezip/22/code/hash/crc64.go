package main

import (
	"flag"
	"fmt"
	"hash/crc64"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Polynomial struct {
	U uint64
}

var polynomials = map[string]uint64{
	"iso":  crc64.ISO,
	"ecma": crc64.ECMA,
}

func (p *Polynomial) Set(s string) error {
	switch s {
	case "iso", "ecma":
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

func (p *Polynomial) Table() *crc64.Table {
	return crc64.MakeTable(p.U)
}

var (
	filename   = flag.String("filename", "crc64.go", "The file to checksum")
	streaming  = flag.Bool("streaming", false, "Whether to stream the file instead of reading it all into memory")
	polynomial = &Polynomial{crc64.ISO}
)

func init() {
	flag.Var(polynomial, "polynomial", "The polynomial to use")
	flag.Parse()
}

func stream(name string) uint64 {
	h := crc64.New(polynomial.Table())
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("failed opening %s: %s", name, err)
	}
	defer file.Close()
	io.Copy(h, file)
	return h.Sum64()
}

func simple(name string) uint64 {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalf("failed reading %s: %s", name, err)
	}
	return crc64.Checksum(data, polynomial.Table())
}

func main() {
	var checksum uint64

	if *streaming {
		checksum = stream(*filename)
	} else {
		checksum = simple(*filename)
	}

	log.Printf("the file %s has checksum %#x", *filename, checksum)
}
