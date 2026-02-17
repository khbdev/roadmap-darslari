package logger

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	path string
}


func NewFileLogger(path string) *FileLogger {
	return &FileLogger{path: path}
}


func (l *FileLogger) Log(msg string) error {
	f, err := os.OpenFile(l.path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	line := fmt.Sprintf("%s %s\n", time.Now().Format(time.RFC3339), msg)
	_, err = f.WriteString(line)
	return err
}