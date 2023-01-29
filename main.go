package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/server"
	"github.com/luozi-csu/lzblogs/utils/logx"
)

var configFile = flag.String("f", "./config.yaml", "path of the global config file")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	w.Write([]byte(message))
}

func main() {
	flag.Parse()

	config.LoadConfig(*configFile)

	logger, err := logx.NewLogger(config.CONF.Server.Logging.Level, config.CONF.Server.Logging.Path)
	if err != nil {
		fmt.Printf("new logger failed, err=%v", err)
		return
	}

	logger.Debugf("this is a debug level message")
	logger.Infof("this is an info level message")
	logger.Warnf("this is a warn level message")

	server, err := server.New(logger)
	if err != nil {
		logger.Errorf("new server failed, err=%v", err)
		return
	}

	err = server.Run()
	if err != nil {
		logger.Errorf("run server failed, err=%v", err)
		return
	}
}
