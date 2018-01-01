package main

import (
	"log"
	"time"
)

var (
	moon = time.Date(1969, time.July, 20, 20, 18, 4, 0, time.UTC)
	now  = time.Now()
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoAdd() {
	log.Println("DemoAdd")
	log.Println(moon.Add(4 * time.Hour))

	log.Println(now)
	// 24 hours from now
	log.Println(now.Add(24 * time.Hour))
	// 24 hours ago, you can add a negative duration
	log.Println(now.Add(-24 * time.Hour))
}

func DemoSub() {
	log.Println("DemoSub")
	log.Println(moon.Sub(time.Now()))
}

func DemoAddDate() {
	log.Println("DemoAddDate")
	log.Println(moon.AddDate(45, 0, 0))
}

func main() {
	log.Println(moon)
	DemoAdd()
	DemoSub()
	DemoAddDate()
}
