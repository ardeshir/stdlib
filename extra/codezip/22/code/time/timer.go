package main

import (
	"log"
	"time"
)

var (
	fiveSeconds = 5 * time.Second
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoTimer() {
	log.Printf("before NewTimer: %s", time.Now())
	t := time.NewTimer(fiveSeconds)
	time.Sleep(3 * time.Second)
	t.Reset(fiveSeconds)
	<-t.C
	// Should be at least 8 seconds later
	log.Printf(" after NewTimer: %s", time.Now())
}

func DemoSleep() {
	log.Printf("before Sleep: %s", time.Now())
	time.Sleep(fiveSeconds)
	// Five seconds later
	log.Printf(" after Sleep: %s", time.Now())
}

func DemoAfter() {
	log.Printf("before After: %s", time.Now())
	now := <-time.After(fiveSeconds)
	// Five seconds later
	log.Printf(" after After: %s", now)
}

func DemoAfterFunc() {
	c := make(chan time.Time)
	log.Printf("before AfterFunc: %s", time.Now())
	time.AfterFunc(fiveSeconds, func() {
		// Otherwise, the program would
		// end without this getting called
		c <- time.Now()
	})
	// Five seconds later
	log.Printf(" after AfterFunc: %s", <-c)
}

func main() {
	DemoTimer()
	DemoSleep()
	DemoAfter()
	DemoAfterFunc()
}
