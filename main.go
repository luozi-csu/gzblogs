package main

import (
	"flag"
	"fmt"

	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/utils/logx"
)

var configFile = flag.String("f", "./config.yaml", "path of the global config file")

func main() {
	flag.Parse()

	config.LoadConfig(*configFile)

	logLevel := config.CONF.Logging.Level
	logPath := config.CONF.Logging.Path
	levelMap := map[string]int{
		"debug": 0, "info": 1, "warn": 2, "error": 3, "fatal": 4,
	}

	logx.InitLogger(levelMap[logLevel], logPath)

	fmt.Println(config.CONF)

	logx.Debugf("this is a debug level message")
	logx.Infof("this is a info level message")
	logx.Warnf("this is a warn level message")
	logx.Errorf("this is a error level message")
	logx.Fatalf("this is a fatal level message")
}
