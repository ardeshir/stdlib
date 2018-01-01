package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

var n int

type Thing struct {
	N int
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	runtime.GOMAXPROCS(8)
}

func Run(pool *sync.Pool) {
	for i := 0; i < 5; i++ {
		thing := pool.Get().(*Thing)
		log.Println(thing.N)
		pool.Put(thing)
	}
}

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			n += 1
			return &Thing{n}
		},
	}

	go Run(pool)
	go Run(pool)
	go Run(pool)
	go Run(pool)
	go Run(pool)

	time.Sleep(time.Second)
}
