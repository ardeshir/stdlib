package main

import (
	"flag"
	"log"
	"runtime/debug"
)

var (
	gcPercent = flag.Int("gc", 100, "garbage collection target percentage")
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func C() {
	debug.PrintStack()
}

func B() {
	C()
}

func A() {
	B()
}

func DemoGCStats() {
	var gc debug.GCStats
	debug.ReadGCStats(&gc)
	log.Printf("LastGC:\t%s", gc.LastGC)
	log.Printf("PauseTotal:\t%s", gc.PauseTotal)
	log.Printf("NumGC:\t%d", gc.NumGC)
	log.Printf("Pause:\t%s", gc.Pause)
}

func main() {
	flag.Parse()
	debug.SetGCPercent(*gcPercent)
	A()
	DemoGCStats()
}
