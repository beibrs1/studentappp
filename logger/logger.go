package logger

import (
	"github.com/fatih/color"
	"time"
)

func Info(msg string) {
	color.New(color.FgGreen).Printf("[INFO] %s — %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func Error(msg string) {
	color.New(color.FgRed).Printf("[ERROR] %s — %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}
