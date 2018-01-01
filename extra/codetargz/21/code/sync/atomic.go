package main

import (
	"log"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	expected int32 = 1000 * 1000
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	runtime.GOMAXPROCS(8)
}

func DemoBroken() {
	var n int32
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				n++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	log.Printf("got %d, expected %d", n, expected)
}

func DemoAtomic() {
	var n int32
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&n, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	log.Printf("got %d, expected %d", n, expected)
}

func main() {
	DemoBroken()
	DemoAtomic()
}
