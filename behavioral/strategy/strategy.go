package strategy

import (
	"fmt"
)

// Abstraction
type LoggerStrategy interface {
	Write(string) error
}

// Strategy 1
type FileLogger struct {}
func (l FileLogger) Write(data string) error {
	fmt.Println("Writed to file...")
	return nil
}

// Strategy 2
type DbLogger struct {}
func (l DbLogger) Write(data string) error {
	fmt.Println("Writed to db..")
}

// Consumer
type Logger struct{
	logger LoggerStrategy
}

func (l *Logger) Log(data string) error {
	l.logger.Write(data)
}


