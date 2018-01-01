package main

import (
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	defer func() {
		logger.SetFlags(log.LstdFlags)
		if err := recover(); err != nil {
			logger.Fatalf("recovered: %s", err)
		}
	}()

	logger.Println("just a string")
	logger.SetPrefix("[go-thestdlib] ")
	logger.Printf("the time is %s", time.Now())
	logger.SetFlags(log.Lshortfile)
	logger.Println("see, the format changed?")
	logger.Panicf("don't worry, we'll handle this")
}
