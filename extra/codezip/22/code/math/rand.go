package main

import (
	"log"
	"math/rand"
)

func example(seed int64) {
	s := rand.NewSource(seed)
	r := rand.New(s)
	log.Printf("ExpFloat64: %f", r.ExpFloat64())
	log.Printf("Float32: %f", r.Float32())
	log.Printf("Float64: %f", r.Float64())
	log.Printf("Int: %d", r.Int())
	log.Printf("Int31: %d", r.Int31())
	log.Printf("Int31n: %d", r.Int31n(10))
	log.Printf("Int63: %d", r.Int63())
	log.Printf("Int63n: %d", r.Int63n(15))
	log.Printf("Intn: %d", r.Intn(25))
	log.Printf("NormFloat64: %f", r.NormFloat64())
	log.Printf("Perm: %v", r.Perm(10))
	log.Printf("Uint32: %d", r.Uint32())
}

func main() {
	example(64)
	// Will print the same as above
	example(64)
	example(1)
}
