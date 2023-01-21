package middleware

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/luozi-csu/lzblogs/util/logx"
)

type logger struct{}

func (l logger) Print(v ...interface{}) {
	logx.Debugf(fmt.Sprint(v...))
}

func RequestLogger(next http.Handler) http.Handler {
	requestLogger := middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logger{},
		NoColor: true,
	})
	return requestLogger(next)
}
