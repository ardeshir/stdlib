package main

import (
	"io"
	"log"
	"os/exec"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoExec() {
	cmd := exec.Command("date", "-u")
	out, err := cmd.Output()
	if err != nil {
		log.Printf("failed running command: %s", err)
	}
	log.Printf("date -u: %s", out)
}

func DemoExecInput() {
	cmd := exec.Command("ruby", "-r", "active_support/all")

	wr, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("failed getting stdin: %s", err)
	}

	rd, _ := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("failed getting stdout: %s", err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("failed starting command: %s", err)
	}
	defer cmd.Wait()

	io.WriteString(wr, "puts 1.hour;")
	io.WriteString(wr, "puts 1.day;")
	wr.Close()

	hour := make([]byte, 5)
	rd.Read(hour)
	log.Printf("1.hour: %s", hour)

	day := make([]byte, 6)
	rd.Read(day)
	log.Printf("1.hour: %s", day)
}

func main() {
	DemoExec()
	DemoExecInput()
}
