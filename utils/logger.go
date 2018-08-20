package utils

import (
	"github.com/caoxiaolin/go-shorturl/config"
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	logFile, err := os.OpenFile(config.Cfg.Log.Logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}
