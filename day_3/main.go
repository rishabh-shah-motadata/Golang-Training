package day3

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	log(msg string)
}

type app struct {
	Logger Logger
}

func (a app) Run(msg string) {
	a.Logger.log(msg)
}

type fileLogger struct {
	Path string
}

func (f fileLogger) log(msg string) {
	file, _ := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(msg + "\n")
}

type consoleLogger struct{}

func (c consoleLogger) log(msg string) {
	fmt.Println(msg)
}

type timestampLogger struct {
	decorator Logger
}

func (t timestampLogger) log(msg string) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	t.decorator.log(ts + " " + msg)
}

func Day3() {
	app1 := app{Logger: timestampLogger{decorator: fileLogger{Path: "day_3/app.log"}}}
	app1.Run("hello from file logger")

	app2 := app{Logger: consoleLogger{}}
	app2.Run("hello from console logger")
}
