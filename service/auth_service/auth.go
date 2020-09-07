package auth_service

import (
	"go-gin-duty-master/models"
)

type Auth struct {
	Id              int
	Name            string
	Telephone       string
	Group           string
	Username        string
	Password        string
	IsAdministrator int
}

func (a *Auth) AddAuth() error {
	data := make(map[string]interface{})
	data["name"] = a.Name
	data["telephone"] = a.Telephone
	data["group"] = a.Group
	data["username"] = a.Username
	data["password"] = a.Password
	data["created_by"] = a.Name
	return models.AddAuth(data)
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

func (a *Auth) GetGroupByName() (string, error) {
	return models.GetGroup(a.Name)
}

func (a *Auth) IsExistName() (bool, error) {
	return models.IsExistName(a.Username, a.Password)
}
