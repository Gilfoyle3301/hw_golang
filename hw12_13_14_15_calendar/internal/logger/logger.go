package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	newLoggerObject *slog.Logger
}

func NewLogger() *Logger {
	newLoggerObject := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &Logger{newLoggerObject}
}

func (l Logger) Info(msg string) {
	l.newLoggerObject.Info(msg)
}

func (l Logger) Error(msg string) {
	l.newLoggerObject = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	l.newLoggerObject.Error(msg)
}

func (l Logger) Warn(msg string) {
	l.newLoggerObject.Warn(msg)
}

func (l Logger) Debug(msg string) {
	l.newLoggerObject.Debug(msg)
}
