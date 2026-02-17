package logger

import "fmt"



func NewLogger() (Logger, error){
	t := "file"
	switch  t {
	case "file":
		path := "app.log"
		if path == "" {
			path = "logger.log"
		}
		return  NewFileLogger(path), nil
	case "db":
		dsn := "postgres://user:pass@localhost:5432/logdb?sslmode=disable"
			if dsn == "" {
			return nil, fmt.Errorf("LOGGER_DB_DSN is empty")
		}
		return NewDBLogger(dsn), nil
			default:
		return nil, fmt.Errorf("unknown LOGGER_TYPE: %q (use file|db)", t)

		
	}

}