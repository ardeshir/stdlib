package main

import (
	"log"
	"time"
)

var (
	utcPlusOne = time.FixedZone("UTC+1", 3600)
	layout     = "Jan _2 15:04:05 2006"
	moon       = "Jul 20 20:18:04 1969"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	log.Println(time.LoadLocation("Canada/Mountain"))

	moonTime, err := time.Parse(layout, moon)
	// Defaults to UTC, kind of wish it defaulted to local
	log.Println(moonTime, err)

	// Same time, different timezone
	moonTime, err = time.ParseInLocation(layout, "Jul 20 21:18:04 1969", utcPlusOne)
	log.Println(moonTime, err)
	log.Println(moonTime.In(time.UTC))

	now := time.Now()
	log.Println(now)
	log.Println(now.In(utcPlusOne))
	log.Println(now.In(time.UTC))
}
