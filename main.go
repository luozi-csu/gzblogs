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

	conf, err := config.NewConfig(*configFile)
	if err != nil {
		fmt.Printf("create config failed, err=%v", err)
		return
	}

	err = logx.Init(&conf.Server.Logging)
	if err != nil {
		fmt.Printf("initialize logger failed, err=%v", err)
		return
	}

	server, err := server.New()
	if err != nil {
		logx.Fatalf("create server failed, err=%v", err)
	}

	if err = server.Run(); err != nil {
		logx.Fatalf("run server failed, err=%v", err)
	}
}
