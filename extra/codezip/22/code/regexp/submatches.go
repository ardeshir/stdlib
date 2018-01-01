package main

import (
	"log"
	"regexp"
)

type Matcher struct {
	*regexp.Regexp
}

func (m *Matcher) FindAllStringSubmatchMap(s string) map[string]string {
	pairs := make(map[string]string)

	// Ignore the first one, it's the "whole" match
	subexpNames := m.SubexpNames()[1:]
	submatches := m.FindAllStringSubmatch(s, -1)
	if submatches == nil {
		return pairs
	}

	// Ignore the first one, it's the "whole" match
	for index, submatch := range submatches[0][1:] {
		name := subexpNames[index]
		pairs[name] = submatch
	}
	return pairs
}

var (
	nginxLogFormat = &Matcher{regexp.MustCompile(`(?P<RemoteAddr>\S+) (?P<Host>\S+) - \[(?P<Time>[^\]]+)\] "(?P<Method>\S+) (?P<Path>\S+) [^"]+" (?P<Status>\d+) (?P<Bytes>\d+) "(?P<UserAgent>[^"]+)" "(?P<Referer>[^"]+)" (?P<RequestTime>\d+\.\d+)`)}
	// log_format timed_combined '$remote_addr $host $remote_user [$time_local] "$request" $status $body_bytes_sent "$http_referer" "$http_user_agent" $request_time';
	logLine = `74.86.158.107 example.com - [01/Dec/2013:18:07:26 -0700] "GET /en/landing HTTP/1.1" 302 108 "-" "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)" 0.087`
)

func main() {
	log.Printf("NumSubexp: %d", nginxLogFormat.NumSubexp())
	subexpNames := nginxLogFormat.SubexpNames()
	log.Printf("SubexpNames: %v", subexpNames)
	submatches := nginxLogFormat.FindAllStringSubmatch(logLine, -1)
	log.Println(submatches)
	log.Println(nginxLogFormat.FindAllStringSubmatchMap(logLine))
}
