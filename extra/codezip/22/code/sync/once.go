package main

import (
	"log"
	"runtime"
	"sync"
)

var once sync.Once

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	runtime.GOMAXPROCS(8)
}

func main() {
	f := func() {
		log.Println("Hello!")
	}
	once.Do(f) // Called
	once.Do(f) // Not called
}
