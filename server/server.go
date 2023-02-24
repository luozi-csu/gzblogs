package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/authentication"
	"github.com/luozi-csu/lzblogs/authorization/oauth"
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
		conf:   conf,
	}

	db, err := database.NewMysql(&conf.Database)
	if err != nil {
		return nil, errors.Wrap(err, "init database failed")
	}

	adapter, err := database.NewMysqlAdapter(db)
	if err != nil {
		return nil, errors.Wrap(err, "create database adapter failed")
	}

	repository := repository.NewRepository(db, conf.Server.RBACModelConf, adapter)
	if conf.Database.Migrate {
		if err = repository.Migrate(); err != nil {
			return nil, errors.Wrap(err, "database auto migration failed")
		}
	}

	// oauth manager
	oauthManager := oauth.NewOAuthManager(&conf.OAuth)

	// services
	userService := service.NewUserService(repository.User())
	rbacService := service.NewRBACService(repository.RBAC())
	jwtService := authentication.NewJWTService(conf.Server.JWTSecret)

	// controllers
	userController := controller.NewUserController(userService)
	rbacController := controller.NewRBACController(rbacService)
	authController := controller.NewAuthController(userService, jwtService, oauthManager)

	s.controllers = append(s.controllers, userController, rbacController, authController)

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
