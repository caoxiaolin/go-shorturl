package utils

import (
	"github.com/caoxiaolin/go-shorturl/config"
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	logFileName := config.Cfg.Log.Logpath + config.Cfg.Log.Logfile
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}
