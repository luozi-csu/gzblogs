package logx

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/luozi-csu/lzblogs/config"
)

const (
	DebugLevel int = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var logger logHelper

type logHelper struct {
	level   *string
	logPath *string
	fd      *os.File
}

func Debugf(format string, args ...interface{}) {
	if logger.getLogLevel() <= DebugLevel {
		log.SetPrefix("[debug] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Infof(format string, args ...interface{}) {
	if logger.getLogLevel() <= InfoLevel {
		log.SetPrefix("[info] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Warnf(format string, args ...interface{}) {
	if logger.getLogLevel() <= WarnLevel {
		log.SetPrefix("[warn] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Errorf(format string, args ...interface{}) {
	if logger.getLogLevel() <= ErrorLevel {
		log.SetPrefix("[error] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Fatalf(format string, args ...interface{}) {
	if logger.getLogLevel() <= FatalLevel {
		log.SetPrefix("[fatal] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger logHelper) Write(buf []byte) (n int, err error) {
	logger.createLogFile()

	if logger.fd == nil {
		fmt.Printf("console: %s", buf)
		return len(buf), nil
	}

	return logger.fd.Write(buf)
}

func (logger *logHelper) createLogFile() {
	logdir := "./"
	logPath := *logger.logPath
	index := strings.LastIndex(logPath, "/")
	if index != -1 {
		logdir = logPath[0:index] + "/"
		os.MkdirAll(logPath[0:index], os.ModePerm)
	}

	now := time.Now()

	var prefix string
	if index != -1 && index != len(logPath)-1 {
		prefix = logPath[index+1:] + "_"
	}

	filename := fmt.Sprintf("%s%04d%02d%02d.log", prefix, now.Year(), now.Month(), now.Day())
	for i := 0; i < 10; i++ {
		if fd, err := os.OpenFile(logdir+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm); err == nil {
			logger.fd.Sync()
			logger.fd.Close()
			logger.fd = fd
			break
		}

		logger.fd = nil
	}
}

func (logger *logHelper) getLogLevel() (level int) {
	logLevelStr := *logger.level
	switch logLevelStr {
	case "debug":
		level = DebugLevel
	case "info":
		level = InfoLevel
	case "warn":
		level = WarnLevel
	case "error":
		level = ErrorLevel
	case "fatal":
		level = FatalLevel
	}
	return
}

func init() {
	logger.level = &config.CONF.Server.Logging.Level
	logger.logPath = &config.CONF.Server.Logging.Path

	log.SetOutput(logger)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
