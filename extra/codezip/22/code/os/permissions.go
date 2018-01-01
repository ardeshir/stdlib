package main

import (
	"flag"
	"log"
	"os"
	"strconv"
)

var (
	chmod = flag.String("chmod", "", "the file to chmod")
	mode  = flag.String("mode", "", "the mode to set")
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	flag.Parse()
}

func main() {
	fileMode, err := strconv.ParseUint(*mode, 8, 32)
	if err != nil {
		log.Fatalf("invalid mode: %s", err)
	}

	err = os.Chmod(*chmod, os.FileMode(fileMode))
	if err != nil {
		log.Fatalf("failed chmod: %s", err)
	}
}
