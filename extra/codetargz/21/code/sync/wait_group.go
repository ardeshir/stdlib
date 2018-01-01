package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

var n = 5

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	runtime.GOMAXPROCS(8)
}

func Run(id int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		wg.Done()
		log.Printf("%d is done", id)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(n)
		go Run(i, &wg)
	}
	wg.Wait()
	log.Println("all done")
}
