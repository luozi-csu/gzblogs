package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/middleware"
	"github.com/luozi-csu/lzblogs/utils/logx"
)

type Server struct {
	e      *gin.Engine
	logger *logx.Logger
}

func New(logger *logx.Logger) (*Server, error) {
	s := &Server{
		e:      gin.Default(),
		logger: logger,
	}

	s.e.Use(middleware.RequestLogger(logger))
	s.e.GET("/", func (c *gin.Context)  {
		c.JSON(200, "hello world")
	})

	return s, nil
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: s.e,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatalf("Failed to start server, %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch := <-sig
	s.logger.Infof("Receive signal: %s", ch)

	return server.Shutdown(ctx)
}
