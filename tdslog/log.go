package tdslog

import (
	"github.com/golang/glog"
	"log"
	"os"
	"sync"
)

const logFileName = "tds.log"

var lock = &sync.Mutex{}
var iLog *log.Logger

func Log(format string, v ...interface{}) {
	get().Printf(format, v...)
}

func get() *log.Logger {
	if iLog == nil {
		defer lock.Unlock()
		lock.Lock()
		if iLog == nil {
			writer, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				glog.Fatal("Error opening log file", err)
				panic(err)
			}

			iLog = log.New(writer, "", log.Ltime)
		}
	}

	return iLog
}
