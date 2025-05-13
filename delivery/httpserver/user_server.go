package myhttpserver

import (
	"fmt"
	"net/http"

	"github.com/fatemehmirarab/gameapp/service/userservice"
	"github.com/labstack/echo/v4"
)

func (s Server) UserRegister(c echo.Context) error {

	var req userservice.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	response, errRegister := s.UserSvc.Register(req)
	if errRegister != nil {
		fmt.Println("❌ Register failed:", errRegister)
		return echo.NewHTTPError(http.StatusBadRequest, errRegister.Error())
	}

	return c.JSON(http.StatusCreated, response)
}

func (s Server) UserLogin(c echo.Context) error {

	var loginReq userservice.LoginRequest

	if err := c.Bind(&loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	response, errRegister := s.UserSvc.Login(loginReq)
	if errRegister != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, response)

}

func (s Server) UserProfile(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")
	claims, tokenErr := s.AuthSvc.ParseJWT(token)

	if tokenErr != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	profile, errProfile := s.UserSvc.Profile(userservice.ProfileRequest{UserId: claims.UserId})
	if errProfile != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, profile)
}
