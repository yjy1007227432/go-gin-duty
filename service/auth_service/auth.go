package auth_service

import (
	"go-gin-duty-master/models"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}

func (a *Auth) GetNameByUsername() (string, error) {
	return models.GetNameByUsername(a.Username)
}

func (a *Auth) IsAdmin() (int, error) {
	return models.IsAdmin(a.Username)
}
