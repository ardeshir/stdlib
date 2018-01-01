package main

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	stop := make(chan bool)
	ticker := time.NewTicker(time.Second)
	time.AfterFunc(5*time.Second, func() {
		ticker.Stop()
		stop <- true
	})

	for {
		select {
		case now := <-ticker.C:
			log.Println(now)
		case <-stop:
			log.Println("stopped")
			return
		}
	}
}
