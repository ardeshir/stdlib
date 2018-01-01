package main

import (
	"log"
	"time"
)

var (
	utcPlusOne = time.FixedZone("UTC+1", 3600)
	moon       = time.Date(1969, time.July, 20, 20, 18, 4, 0, time.UTC)
	moonAlso   = time.Date(1969, time.July, 20, 21, 18, 4, 0, utcPlusOne)
	now        = time.Now()
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoBefore() {
	log.Println("DemoBefore")
	log.Printf("moon before now? %t", moon.Before(now))
}

func DemoAfter() {
	log.Println("DemoAfter")
	log.Printf("moon after now? %t", moon.After(now))
}

func DemoEqual() {
	log.Println("DemoEqual")
	log.Printf("moon equal now? %t", moon.Equal(now))
	log.Printf("moon equal moon? %t", moon.Equal(moon))

	log.Printf("moon: %s", moon)
	log.Printf("moonAlso: %s", moonAlso)
	log.Printf("moon equal moonAlso? %t", moon.Equal(moonAlso))
}

func main() {
	DemoBefore()
	DemoAfter()
	DemoEqual()
}
