package main

import (
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	signals = make(chan os.Signal, 1)
	val     int32
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")

	signal.Notify(signals, syscall.SIGUSR1, syscall.SIGUSR2)
	go handleSignals()
}

func handleSignals() {
	for signal := range signals {
		switch signal {
		case syscall.SIGUSR1:
			log.Println("got USR1, adding 2")
			atomic.AddInt32(&val, 2)
		case syscall.SIGUSR2:
			log.Println("got USR2, subtracting 1")
			atomic.AddInt32(&val, -1)
		}
		log.Printf("val: %d", val)
	}
}

func main() {
	os.Getpid()
	p, _ := os.FindProcess(os.Getpid())

	ticker := time.Tick(1 * time.Second)
	for now := range ticker {
		switch {
		case val > 5:
			p.Kill()
		case now.Second()%2 == 0: // even
			p.Signal(syscall.SIGUSR1)
		case now.Second()%2 == 1: // odd
			p.Signal(syscall.SIGUSR2)
		}
	}
}
