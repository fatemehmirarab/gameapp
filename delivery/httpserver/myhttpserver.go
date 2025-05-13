package myhttpserver

import (
	"fmt"
	"log"

	"github.com/fatemehmirarab/gameapp/config"
	"github.com/fatemehmirarab/gameapp/service/userservice"
	"github.com/fatemehmirarab/gameapp/service/userservice/authservice"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Config  config.Config
	AuthSvc authservice.Service
	UserSvc userservice.Service
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service) Server {
	return Server{
		Config:  config,
		AuthSvc: authSvc,
		UserSvc: userSvc,
	}
}

func (s Server) Serve() {
	log.Println("start")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	userGroup := e.Group("/user")
	userGroup.POST("/register", s.UserRegister)
	userGroup.POST("/login", s.UserLogin)
	userGroup.GET("/profile", s.UserProfile)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%d", s.Config.HttpServer.Port)))

}
