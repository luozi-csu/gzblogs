package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	debugLevel int = iota
	infoLevel
	warnLevel
	errorLevel
	fatalLevel
)

var Logger logHelper

type logHelper struct {
	level   int
	logPath string
	logTime int64
	fd      *os.File
}

func InitLogger(level int, path string) {
	Logger.level = level
	Logger.logPath = path
	Logger.logTime = time.Now().Unix()
	Logger.createLogFile()

	log.SetOutput(Logger)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (logger *logHelper) SetLevel(level int) {
	logger.level = level
}

func (logger logHelper) Debugf(format string, args ...interface{}) {
	if logger.level >= debugLevel {
		log.SetPrefix("[debug] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger logHelper) Infof(format string, args ...interface{}) {
	if logger.level >= infoLevel {
		log.SetPrefix("[info] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger logHelper) Warnf(format string, args ...interface{}) {
	if logger.level >= warnLevel {
		log.SetPrefix("[warn] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger logHelper) Errorf(format string, args ...interface{}) {
	if logger.level >= errorLevel {
		log.SetPrefix("[error] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger logHelper) Fatalf(format string, args ...interface{}) {
	if logger.level >= fatalLevel {
		log.SetPrefix("[fatal] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger logHelper) Write(buf []byte) (n int, err error) {
	if logger.fd == nil {
		fmt.Printf("console: %s", buf)
		return len(buf), nil
	}

	var daySeconds int64 = 24 * 60 * 60
	if logger.logTime+daySeconds < time.Now().Unix() {
		logger.createLogFile()
		logger.logTime = time.Now().Unix()
	}

	return logger.fd.Write(buf)
}

func (logger *logHelper) createLogFile() {
	logdir := "./"
	index := strings.LastIndex(logger.logPath, "/")
	if index != -1 {
		logdir = logger.logPath[0:index] + "/"
		os.MkdirAll(logger.logPath[0:index], os.ModePerm)
	}

	now := time.Now()

	var prefix string
	if index != -1 && index != len(logger.logPath)-1 {
		prefix = logger.logPath[index+1:] + "_"
	}

	filename := fmt.Sprintf("%s%04d%02d%02d.log", prefix, now.Year(), now.Month(), now.Day())
	for i := 0; i < 10; i++ {
		if fd, err := os.OpenFile(logdir+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm); err == nil {
			logger.fd.Sync()
			logger.fd.Close()
			logger.fd = fd
			break
		} else {
			fmt.Println(err)
		}

		logger.fd = nil
	}
}
