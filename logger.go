package main

import (
	"log"
	"os"
	"io/ioutil"
)

type Logger struct {
	Error   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Trace   *log.Logger
}

func NewLogger(level string) *Logger {
	return &Logger{
		Error:   log.New(os.Stdout, "[Error]", log.LstdFlags),
		Warning: log.New(os.Stdout, "[Warning]", log.LstdFlags),
		Info:    log.New(os.Stderr, "[Info]", log.LstdFlags),
		Trace:   log.New(ioutil.Discard, "[Trace]", log.LstdFlags),
	}
}

