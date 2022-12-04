package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/middleware"
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

	logLevel := config.CONF.Logging.Level
	logPath := config.CONF.Logging.Path
	levelMap := map[string]int{
		"debug": 0, "info": 1, "warn": 2, "error": 3, "fatal": 4,
	}

	logx.InitLogger(levelMap[logLevel], logPath)

	router := chi.NewRouter()
	router.Use(middleware.RequestLogger)

	router.Get("/", helloHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		logx.Errorf("listen and serve http failed, err=%v", err)
		return
	}
}
