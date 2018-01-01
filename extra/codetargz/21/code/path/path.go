package main

import (
	"flag"
	"log"
	"path"
)

func main() {
	var p string
	flag.StringVar(&p, "path", "./foo/../baz.gif", "the path to examine")
	flag.Parse()

	log.Printf("p: %s", p)
	log.Printf("Base(p): %s", path.Base(p))
	log.Printf("Clean(p): %s", path.Clean(p))
	log.Printf("Dir(p): %s", path.Dir(p))
	log.Printf("Ext(p): %s", path.Ext(p))
	log.Printf("IsAbs(p): %t", path.IsAbs(p))
	log.Printf("Join(\"/fizz/bin\", p): %s", path.Join("/fizz/bin", p))

	matched, err := path.Match("/*/bin/*.gif", p)
	log.Printf("Match(\"/*/bin/*.gif\", p): %t, %v", matched, err)

	matched, err = path.Match("/*/bin/*.gif", path.Join("/fizz/bin", p))
	log.Printf("Match(\"/*/bin/*.gif\", Join(\"/fizz/bin\", p)): %t, %v", matched, err)

	dir, file := path.Split(p)
	log.Printf("Split(p): %s, %s", dir, file)
}
