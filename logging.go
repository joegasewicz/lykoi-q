package main

import (
	"fmt"
	"log"
	"time"
)

const (
	DEBUG = "DEBUG"
	PROD  = "PROD"
)

type Logging struct {
	Level   string
	Threads uint8
	Port    int
}

func NewLogging(level string, threads uint8) *Logging {
	if level == "" {
		level = PROD
	}
	return &Logging{
		Level:   level,
		Threads: threads,
	}
}

func (l *Logging) Log(msg ...string) {
	currentTime := time.Now()
	logOut := fmt.Sprintf(
		"[%v: %v/threads=%v] ",
		currentTime.Format("2017-09-07 17:06:06"),
		l.Level,
		l.Threads,
	)
	for i := 0; i < len(msg); i++ {
		if i == len(msg)-1 {
			logOut += msg[i]
		} else {
			logOut += msg[i] + " "
		}

	}
	log.SetFlags(0)
	log.Println(logOut)
}
