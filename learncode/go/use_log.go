package main

import (
	"log"
	"os"
)

const _LOG_DEBUG_FLAG = log.Ldate | log.Lmicroseconds | log.Lshortfile

func main() {
	var errorlog *log.Logger
	if fd, err := os.OpenFile("./chenp", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		errorlog = log.New(fd, "epassport: ", _LOG_DEBUG_FLAG)
	} else {
		return
	}
	errorlog.Println("shit")

}
