package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/middleware"
	"github.com/luozi-csu/lzblogs/util/logx"
)

var configFile = flag.String("f", "./config.yaml", "path of the global config file")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	w.Write([]byte(message))
}

func main() {
	flag.Parse()

	config.LoadConfig(*configFile)

	logx.Debugf("this is a debug level message")
	logx.Infof("this is an info level message")

	router := chi.NewRouter()
	router.Use(middleware.RequestLogger)

	router.Get("/", helloHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.CONF.Server.Port), router)
	if err != nil {
		logx.Errorf("listen and serve http failed, err=%v", err)
		return
	}
}
