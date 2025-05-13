package config

import (
	mySQL "github.com/fatemehmirarab/gameapp/repository/mysql"
	"github.com/fatemehmirarab/gameapp/service/userservice/authservice"
)

type HttpServer struct {
	Port int
}

type Config struct {
	Auth       authservice.Config
	HttpServer HttpServer
	MySQL      mySQL.Config
}
