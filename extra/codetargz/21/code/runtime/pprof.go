package main

import (
	"log"
	"os"
	"runtime/pprof"
)

const Flags = os.O_CREATE | os.O_TRUNC | os.O_WRONLY

func DumpHeap(name string) {
	file, err := os.OpenFile(name, Flags, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	pprof.Lookup("heap").WriteTo(file, 0)
}

func main() {
	file, err := os.OpenFile("cpu.prof", Flags, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	err = pprof.StartCPUProfile(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer pprof.StopCPUProfile()

	DumpHeap("before.heap")

	fib := []int{0, 1}
	for i := 0; i < 1000000; i++ {
		fib = append(fib, fib[i]+fib[i+1])
	}

	DumpHeap("after.heap")
}
