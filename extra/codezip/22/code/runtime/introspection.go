package main

import (
	"log"
	"runtime"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	log.Printf("GOOS:\t%s", runtime.GOOS)
	log.Printf("GOARCH:\t%s", runtime.GOARCH)
	log.Printf("GOROOT:\t%s", runtime.GOROOT())
	log.Printf("Compiler:\t%s", runtime.Compiler)
	log.Printf("Version:\t%s", runtime.Version())
}
