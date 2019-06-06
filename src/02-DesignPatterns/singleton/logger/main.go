package main

import (
	"log"
	"os"
	"sync"
)

type CustomLogger struct {
	*log.Logger
	filename string
}

var (
	clogger *CustomLogger
	once    sync.Once
)

func GetInstance() *CustomLogger {
	once.Do(func() {
		clogger = createLogger("c.log")
	})
	return clogger
}

func createLogger(filename string) *CustomLogger {
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &CustomLogger{
		filename: filename,
		Logger:   log.New(file, "custom:", log.Lshortfile),
	}
}

func main() {
	logger := GetInstance()
	logger.Println("hello")
}
