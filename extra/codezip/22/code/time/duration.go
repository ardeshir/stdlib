package main

import (
	"log"
	"time"
)

var (
	moon = time.Date(1969, time.July, 20, 20, 18, 4, 0, time.UTC)
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoConstants() {
	log.Println("DemoConstants")
	log.Println(5 * time.Nanosecond)
	log.Println(5 * time.Microsecond)
	log.Println(5 * time.Millisecond)
	log.Println(5 * time.Second)
	log.Println(5 * time.Minute)
	log.Println(5 * time.Hour)
}

func DemoParsing() {
	log.Println("DemoParsing")
	d, _ := time.ParseDuration("5h2m55s10us5ns")
	log.Println(d)
	log.Printf("%fh == %fm == %fs", d.Hours(), d.Minutes(), d.Seconds())
}

func DemoRound() {
	log.Println("DemoRound")
	log.Println(moon)
	log.Println(moon.Round(time.Minute))
	log.Println(moon.Round(time.Hour))
}

func DemoTruncate() {
	// Ignore this math until the next demo
	laterMoon := moon.Add(30 * time.Minute)
	log.Println("DemoTruncate")
	log.Println(laterMoon)
	log.Println(laterMoon.Truncate(time.Hour))
	// See how Round goes up and Truncate goes down?
	log.Println(laterMoon.Round(time.Hour))
}

func DemoSince() {
	log.Println("DemoSince")
	log.Printf("%s since %s", time.Since(moon), moon)
}

func main() {
	DemoConstants()
	DemoParsing()
	DemoRound()
	DemoTruncate()
	DemoSince()
}
