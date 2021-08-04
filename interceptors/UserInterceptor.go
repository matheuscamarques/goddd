package interceptors

import (
	"api/models"
	"api/repository"
)

type UserInterceptor struct {
	repository.UserRepository
}

// valiate user alredy exists
func (ui *UserInterceptor) Exists(user models.User) bool {
	return ui.UserRepository.Exists(user)
}
