package main

import (
	"log"
	"regexp"
	"strings"
)

var (
	redaction = regexp.MustCompile(`(password|token)=(\w+)`)
	pairs     = regexp.MustCompile(`(\w+)=`)
	logLine   = `2013-12-02T02:40:57.049407+00:00 app: at=info method=POST path=/login token=secret host=example.com password=sekrit connect=1ms service=82ms status=200 bytes=809`
)

func main() {
	log.Println(redaction.ReplaceAllString(logLine, "$1=[REDACTED]"))
	log.Println(pairs.ReplaceAllStringFunc(logLine, strings.ToUpper))
}
