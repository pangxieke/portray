package log

import (
	"log"
	"os"
	//"github.com/sirupsen/logrus"
	//"github.com/pangxieke/portray/config"
)

var (
	Log *log.Logger
)

func init() {
	Log = log.New(os.Stderr, "", 1)
}

func Init(logFile string) (err error) {
	lf, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	Log.SetOutput(lf)
	return
}

func Info(args ...interface{}) {
	Log.Print(args...)
}

func Print(args ...interface{}) {
	Log.Print(args...)
}
