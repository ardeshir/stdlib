package main

import (
	"log"
	"os/user"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoCurrent() {
	u, _ := user.Current()
	log.Printf("%#v", u)
}

func DemoLookup() {
	u, _ := user.Lookup("nobody")
	log.Printf("%#v", u)
}

func DemoLookupId() {
	u, _ := user.LookupId("1")
	log.Printf("%#v", u)
}

func main() {
	DemoCurrent()
	DemoLookup()
	DemoLookupId()
}
