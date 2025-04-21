package userservice

import (
	"fmt"

	"github.com/fatemehmirarab/gameapp/entity"
	"github.com/fatemehmirarab/gameapp/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
}

type Service struct {
	Repo Repository
}

func New(repo Repository) Service {
	return Service{Repo: repo}
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
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
	user := entity.User{
		Id:          0,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
	}
	createdUser, err := s.Repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	return RegisterResponse{
		User: createdUser}, err
}
