package main

import (
	"time"

	"github.com/fatemehmirarab/gameapp/config"
	"github.com/fatemehmirarab/gameapp/delivery/myhttpserver"
	mySQL "github.com/fatemehmirarab/gameapp/repository/mysql"
	"github.com/fatemehmirarab/gameapp/service/userservice"
	"github.com/fatemehmirarab/gameapp/service/userservice/authservice"
)

const (
	expirationTime = time.Hour * 24
	accessSubject  = "at"
	refreshSubject = "rt"
)

func main() {

	cfg := config.Config{
		Auth: authservice.Config{
			ExpirationTime:        expirationTime,
			RefreshExpirationTime: expirationTime * 7,
			AccessSubject:         accessSubject,
			RefreshSubject:        refreshSubject,
		},
		HttpServer: config.HttpServer{
			Port: 8080,
		},
		MySQL: mySQL.Config{
			UserName: "fatemeh",
			PassWord: "Mirarab",
			Port:     3306,
			Host:     "localhost",
			DBName:   "mydb",
		},
	}

	authSvc := authservice.New(cfg.Auth)
	mysqlrepo := mySQL.New(cfg.MySQL)
	userSvc := userservice.New(authSvc, mysqlrepo)

	server := myhttpserver.New(cfg, authSvc, userSvc)

	server.Serve()

}
