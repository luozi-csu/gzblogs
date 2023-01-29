package logx

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/luozi-csu/lzblogs/utils"
)

const (
	DebugLevel int = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type Logger struct {
	level   int
	logTime int64
	logPath string
	fd      *os.File
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	if logger.level <= DebugLevel {
		log.SetOutput(logger)
		log.SetPrefix("[debug] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	if logger.level <= InfoLevel {
		log.SetOutput(logger)
		log.SetPrefix("[info] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	if logger.level <= WarnLevel {
		log.SetOutput(logger)
		log.SetPrefix("[warn] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.level <= ErrorLevel {
		log.SetOutput(logger)
		log.SetPrefix("[error] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	if logger.level <= FatalLevel {
		log.SetOutput(logger)
		log.SetPrefix("[fatal] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (logger *Logger) Write(buf []byte) (n int, err error) {
	if logger.fd == nil {
		fmt.Printf("console: %s", buf)
		return len(buf), nil
	}

	if logger.logTime+86400 < time.Now().Unix() {
		logger.createLogFile()
		logger.logTime = utils.Zerotime()
	}

	return logger.fd.Write(buf)
}

func (logger *Logger) createLogFile() {
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
		}

		logger.fd = nil
	}
}

func ParseLevel(levelStr string) (int, error) {
	var level int
	switch strings.ToLower(levelStr) {
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
	default:
		return level, fmt.Errorf("%s is not a valid level", levelStr)
	}

	return level, nil
}

func NewLogger(level, path string) (*Logger, error) {
	lvl, err := ParseLevel(level)
	if err != nil {
		return nil, err
	}
	logger := &Logger{
		level: lvl,
		logPath: path,
		logTime: utils.Zerotime(),
	}
	logger.createLogFile()

	log.SetOutput(logger)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return logger, nil
}
