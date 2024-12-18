package logging

import (
	"log"
)

type Logger struct {
	level string
}

func New(level string) *Logger {
	return &Logger{level: level}
}

func (l *Logger) logMessage(level, message string) {
	if l.shouldLog(level) {
		log.Printf("[%s] %s\n", level, message)
	}
}

func (l *Logger) shouldLog(level string) bool {
	levels := map[string]int{"debug": 1, "info": 2, "warn": 3, "error": 4}
	return levels[level] >= levels[l.level]
}

func (l *Logger) Debug(msg string) {
	l.logMessage("debug", msg)
}

func (l *Logger) Info(msg string) {
	l.logMessage("info", msg)
}

func (l *Logger) Warn(msg string) {
	l.logMessage("warn", msg)
}

func (l *Logger) Error(msg string) {
	l.logMessage("error", msg)
}
