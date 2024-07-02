package logger

import (
	"fmt"
	"log"
)

var logger = log.Default()

func Errorf(format string, v ...any) {
	logger.Printf("[Error]: %s\n", fmt.Sprintf(format, v...))
}

func Infof(format string, v ...any) {
	logger.Printf("[Info]: %s\n", fmt.Sprintf(format, v...))
}
