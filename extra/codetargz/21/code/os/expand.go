package main

import (
	"flag"
	"log"
	"os"
)

type expandable map[string]string

func (e expandable) Expand(s string) string {
	return e[s]
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	flag.Parse()
}

func DemoExpandEnv() {
	log.Println(os.ExpandEnv("$HOME"))
	log.Println(os.ExpandEnv("$PWD"))
}

func DemoExpand() {
	exp := expandable(map[string]string{
		"name":  "Superman",
		"alter": "Clark Kent",
	})
	log.Println(os.Expand("${name} is really ${alter}", exp.Expand))
}

func main() {
	DemoExpandEnv()
	DemoExpand()
}
