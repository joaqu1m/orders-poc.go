package services

import (
	"orders-api/config"
	"orders-api/dtos"
	"orders-api/repositories"
)

func GetUsers() ([]string, error) {
	return repositories.GetUsers()
}

func CreateUser(user dtos.UserRequest) error {
	return repositories.CreateUser(user)
}

func AuthUser(user dtos.UserRequest) (string, error) {

	result, err := repositories.AuthUser(user)
	if err != nil {
		return "", err
	}
	if !result {
		return "", nil
	}

	token, err := config.GenerateToken(user.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}
