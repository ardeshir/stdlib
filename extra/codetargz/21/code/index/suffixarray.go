package main

import (
	"bytes"
	"flag"
	"index/suffixarray"
	"io/ioutil"
	"log"
	T "testing"
)

type Searcher struct {
	index *suffixarray.Index
	n, h  []byte
}

func NewSearcher(n, h []byte) *Searcher {
	return &Searcher{
		n:     n,
		h:     h,
		index: suffixarray.New(h),
	}
}

func (s *Searcher) SuffixarrayPrebuilt() int {
	results := s.index.Lookup(s.n, 1)
	if len(results) == 1 {
		return results[0]
	}
	return -1
}

func (s *Searcher) Suffixarray() int {
	index := suffixarray.New(s.h)
	results := index.Lookup(s.n, 1)
	if len(results) == 1 {
		return results[0]
	}
	return -1
}

func (s *Searcher) BytesIndex() int {
	return bytes.Index(s.h, s.n)
}

var (
	needle   = flag.String("needle", "O Romeo, Romeo! wherefore art thou Romeo?", "The string to search for")
	haystack = flag.String("haystack", "romeo-and-juliet.txt", "The file to search through")
)

func bench(name string, f func() int) {
	index := 0
	result := T.Benchmark(func(b *T.B) {
		for i := 0; i < b.N; i++ {
			index = f()
		}
	})
	log.Printf("%s took %d ns/op to find %#v at index %d", name, result.NsPerOp(), *needle, index)
}

func main() {
	flag.Parse()
	h, err := ioutil.ReadFile(*haystack)
	if err != nil {
		log.Fatalf("failed to read haystack: %s", err)
	}
	s := NewSearcher([]byte(*needle), h)
	bench("SuffixarrayPrebuilt", s.SuffixarrayPrebuilt)
	bench("Suffixarray", s.Suffixarray)
	bench("BytesIndex", s.BytesIndex)
}
