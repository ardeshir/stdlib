package main

import (
	"fmt"
	"log"
	"runtime"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

type movie struct {
	Title string
}

func (m *movie) String() string {
	return fmt.Sprintf("Movie{%s}", m.Title)
}

func DemoFinalizers() {
	logging := make(chan string)

	rockOfAges := &movie{"Rock of Ages"}
	runtime.SetFinalizer(rockOfAges, func(m *movie) {
		logging <- fmt.Sprintf("%s is being cleaned up", m)
		close(logging)
	})

	rockOfAges = nil
	runtime.GC() // Force a GC so the finalizer runs

	for msg := range logging {
		log.Println(msg)
	}
}

func DemoMemstats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc:\t%db", ms.Alloc)
	log.Printf("TotalAlloc:\t%db", ms.TotalAlloc)
	log.Printf("Mallocs:\t%d", ms.Mallocs)
	log.Printf("Frees:\t%d", ms.Frees)
	log.Printf("PauseTotalNs:\t%dns", ms.PauseTotalNs)
}

func main() {
	DemoFinalizers()
	DemoMemstats()
}
