package main

import (
	"factory-1/logger"
	"log"
)

func main(){

		lg, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	_ = lg.Log("server end")

}