package logx

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/utils"
)

const (
	DebugLevel int = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var (
	defaultLogger *logger
)

type logger struct {
	level   int
	logTime int64
	logPath string
	fd      *os.File
}

func Debugf(format string, args ...interface{}) {
	if defaultLogger.level <= DebugLevel {
		log.SetPrefix("[debug] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Infof(format string, args ...interface{}) {
	if defaultLogger.level <= InfoLevel {
		log.SetPrefix("[info] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Warnf(format string, args ...interface{}) {
	if defaultLogger.level <= WarnLevel {
		log.SetPrefix("[warn] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Errorf(format string, args ...interface{}) {
	if defaultLogger.level <= ErrorLevel {
		log.SetPrefix("[error] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Fatalf(format string, args ...interface{}) {
	if defaultLogger.level <= FatalLevel {
		log.SetPrefix("[fatal] ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
	os.Exit(1)
}

func (logger *logger) Write(buf []byte) (n int, err error) {
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

func (logger *logger) createLogFile() {
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

func ParseLevel(lvl string) (int, error) {
	var level int
	switch strings.ToLower(lvl) {
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
		return level, fmt.Errorf("%s is not a valid level", lvl)
	}

	return level, nil
}

func Init(cfg *config.ServerLoggingConfig) error {
	l, err := ParseLevel(cfg.Level)
	if err != nil {
		return err
	}

	defaultLogger = &logger{
		level:   l,
		logPath: cfg.Path,
		logTime: utils.Zerotime(),
	}
	if cfg.IsFile {
		defaultLogger.createLogFile()
	}

	log.SetOutput(defaultLogger)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return nil
}
