package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/controller"
	"github.com/luozi-csu/lzblogs/database"
	"github.com/luozi-csu/lzblogs/middleware"
	"github.com/luozi-csu/lzblogs/repository"
	"github.com/luozi-csu/lzblogs/service"
	"github.com/luozi-csu/lzblogs/utils/logx"
	"github.com/pkg/errors"
)

type Server struct {
	engine      *gin.Engine
	conf        *config.Config
	controllers []controller.Controller
}

func New(conf *config.Config) (*Server, error) {
	s := &Server{
		engine: gin.New(),
	}

	db, err := database.NewMysql(&conf.Database)
	if err != nil {
		return nil, errors.Wrap(err, "init database failed")
	}

	repository := repository.NewRepository(db)

	userService := service.NewUserService(repository.User())
	userController := controller.NewUserController(userService)

	s.controllers = append(s.controllers, userController)

	s.engine.Use(
		gin.Recovery(),
		middleware.RequestLogger,
	)

	return s, nil
}

func (s *Server) Run() error {
	api := s.engine.Group("/api/v1")
	for _, controller := range s.controllers {
		controller.RegisterRoute(api)
	}

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: s.engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logx.Fatalf("Failed to start server, %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch := <-sig
	logx.Infof("Receive signal: %s", ch)

	return server.Shutdown(ctx)
}
