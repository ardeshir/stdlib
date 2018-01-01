package main

import (
	"log"
	"runtime"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	log.Printf("GOMAXPROCS: %d", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use the whole CPU
	log.Printf("GOMAXPROCS: %d", runtime.GOMAXPROCS(0))

	log.Printf("There are %d goroutines running", runtime.NumGoroutine())

	done := make(chan bool)
	go func() {
		log.Println("in the goroutine")

		runtime.LockOSThread()
		log.Println("locked to this OS thread")
		runtime.Gosched() // Let the CPU go

		runtime.UnlockOSThread()
		log.Println("unlocked")
		runtime.Gosched() // Let the CPU go

		// runtime.Goexit() // Will cause a deadlock

		done <- true
	}()

	log.Printf("There are %d goroutines running", runtime.NumGoroutine())
	<-done
}
