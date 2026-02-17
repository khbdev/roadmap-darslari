package logger

import "fmt"

type DBLogger struct {
	dsn string
}

func NewDBLogger(dsn string) *DBLogger {
	return &DBLogger{dsn: dsn}
}

func (l *DBLogger) Log(msg string) error {
	// Hozircha DBga yozish o‘rniga console’ga chiqaryapmiz.
	// Keyin db.Exec("insert into logs...") qilasan.
	fmt.Println("[DB]", msg, "dsn:", l.dsn)
	return nil
}
