package app

import (
	"api/interceptors"
	"api/models"
	"api/repository"
	"errors"
)

type RegisterUser struct {
	repository.UserRepository
	interceptors.UserInterceptor
	User models.User
}

// constructor RegisterUser
func NewRegisterUser(user models.User) *RegisterUser {
	return &RegisterUser{User: user}
}

func (u *RegisterUser) Execute() error {
	exists := u.UserInterceptor.Exists(u.User)
	if exists {
		// return error
		return errors.New("codigo_de_erro")
	}
	return u.UserRepository.Create(u.User)
}
