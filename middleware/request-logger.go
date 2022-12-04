package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

type logger struct{}

func (l logger) Print(v ...interface{}) {
	log.Output(2, fmt.Sprint(v...))
}

func RequestLogger(next http.Handler) http.Handler {
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logger{},
		NoColor: true,
	})
	return middleware.Logger(next)
}
