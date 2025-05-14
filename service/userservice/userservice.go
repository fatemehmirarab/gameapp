package userservice

import (
	"fmt"
	"net/http"

	"crypto/md5"
	"encoding/hex"

	"github.com/fatemehmirarab/gameapp/entity"
	"github.com/fatemehmirarab/gameapp/pkg/errormessage"
	"github.com/fatemehmirarab/gameapp/pkg/phonenumber"
	"github.com/fatemehmirarab/gameapp/pkg/richerror"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserById(userId uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	RefreshToken(user entity.User) (string, error)
}

type Service struct {
	Auth AuthGenerator
	Repo Repository
}

func New(auth AuthGenerator, repo Repository) Service {
	return Service{Auth: auth, Repo: repo}
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
	Password    string
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(request RegisterRequest) (RegisterResponse, error) {
	if !phonenumber.IsValid(request.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	if isUnique, err := s.Repo.IsPhoneNumberUnique(request.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error : %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("Phone Number is not Unique")
		}
	}

	if len(request.Name) <= 3 {
		return RegisterResponse{}, fmt.Errorf("name should be greater than 3")
	}

	if len(request.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password should be greater than 8")
	}

	user := entity.User{
		Id:          0,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Password:    getMD5Hash(request.Password),
	}
	createdUser, err := s.Repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	return RegisterResponse{
		User: createdUser}, err
}

type LoginRequest struct {
	PhoneNumber string
	Password    string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	const op richerror.Op = "userservice.Login"
	user, exist, err := s.Repo.GetUserByPhoneNumber(req.Password)
	if err != nil {
		return LoginResponse{}, richerror.New(op).WithError(err).WithKind(http.StatusInternalServerError).WithMessage(errormessage.UnExpectedError)
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("user or password dose not exist")
	}

	if user.Password != getMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("user or password dose not exist")
	}

	accessToken, err := s.Auth.CreateAccessToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error : %w", err)
	}

	refreshToken, err := s.Auth.RefreshToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	return LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type ProfileRequest struct {
	UserId uint
}
type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	const op richerror.Op = "userservice.Profile"
	user, err := s.Repo.GetUserById(req.UserId)
	if err != nil {
		return ProfileResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"req": req})
	}
	return ProfileResponse{user.Name}, nil

}
