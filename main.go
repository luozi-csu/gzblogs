package main

import (
	"flag"
	"fmt"

	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/server"
	"github.com/luozi-csu/lzblogs/utils/logx"
)

var configFile = flag.String("f", "./config.yaml", "path of the global config file")

func main() {
	flag.Parse()

	config.LoadConfig(*configFile)

	err := logx.ConfigLogger(&config.CONF.Server.Logging)
	if err != nil {
		fmt.Printf("initialize logger failed, err=%v", err)
		return
	}

	logx.Debugf("this is a debug level message")
	logx.Infof("this is an info level message")
	logx.Warnf("this is a warn level message")

	server, err := server.New()
	if err != nil {
		logx.Fatalf("create server failed, err=%v", err)
		return
	}

	if err = server.Run(); err != nil {
		logx.Fatalf("run server failed, err=%v", err)
		return
	}
}
