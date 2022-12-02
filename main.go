package main

import (
	"flag"
	"fmt"

	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/utils"
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

	utils.InitLogger(levelMap[logLevel], logPath)

	fmt.Println(config.CONF)

	utils.Logger.Debugf("this is a debug level message")
	utils.Logger.Infof("this is a info level message")
	utils.Logger.Warnf("this is a warn level message")
	utils.Logger.Errorf("this is a error level message")
	utils.Logger.Fatalf("this is a fatal level message")
}
