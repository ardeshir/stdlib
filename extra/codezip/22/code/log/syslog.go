package main

import (
	"log"
	"log/syslog"
)

func MustSyslog(p syslog.Priority, flags int) *log.Logger {
	logger, err := syslog.NewLogger(p, flags)
	if err != nil {
		panic(err)
	}
	return logger
}

type Logger struct {
	Alert, Crit, Debug, Emerg, Err, Info, Notice, Warning *log.Logger
}

func NewLogger(flags int) *Logger {
	return &Logger{
		Alert:   MustSyslog(syslog.LOG_ALERT, flags),
		Crit:    MustSyslog(syslog.LOG_CRIT, flags),
		Debug:   MustSyslog(syslog.LOG_DEBUG, flags),
		Emerg:   MustSyslog(syslog.LOG_EMERG, flags),
		Err:     MustSyslog(syslog.LOG_ERR, flags),
		Info:    MustSyslog(syslog.LOG_INFO, flags),
		Notice:  MustSyslog(syslog.LOG_NOTICE, flags),
		Warning: MustSyslog(syslog.LOG_WARNING, flags),
	}
}

func basic() {
	logger, err := syslog.NewLogger(syslog.LOG_WARNING, log.Lshortfile)
	if err != nil {
		log.Fatalf("failed to make syslogger: %s", err)
	}
	logger.Println("hello, world")
}

func levels() {
	logger := NewLogger(log.Lshortfile)
	logger.Crit.Println("oh noes!")
	logger.Warning.Println("just a warning")
	logger.Alert.Println("alert time!")
}

func writer() {
	logger, err := syslog.New(syslog.LOG_WARNING, "go-thestdlib")
	if err != nil {
		log.Fatalf("failed to make a syslogger: %s", err)
	}
	logger.Warning("just a message...")
}

func main() {
	basic()
	levels()
	writer()
}
