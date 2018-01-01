package main

import (
	"flag"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	flag.Parse()
}

func DemoProcessIds() {
	log.Printf("process id: %d", os.Getpid())
	log.Printf("parent process id: %d", os.Getppid())
}

func DemoUserInfo() {
	// actually running the program
	log.Printf("user id: %d", os.Getuid())
	log.Printf("group id: %d", os.Getgid())

	// permissions
	log.Printf("effective user id: %d", os.Geteuid())
	log.Printf("effective group id: %d", os.Getegid())

	groups, err := os.Getgroups()
	if err != nil {
		log.Fatalf("failed getting groups: %s", err)
	}
	log.Printf("groups you belong to: %d", groups)
}

func DemoExtra() {
	log.Printf("$GOPATH: %s", os.Getenv("GOPATH"))
	log.Printf("$TMPDIR: %s", os.Getenv("TMPDIR"))

	log.Printf("pagesize: %d bytes", os.Getpagesize())

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("failed getting hostname: %s", err)
	}
	log.Printf("hostname: %s", hostname)
}

func main() {
	DemoProcessIds()
	DemoUserInfo()
	DemoExtra()
}
