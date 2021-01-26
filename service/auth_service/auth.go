package auth_service

import (
	"go-gin-duty-master/models"
)

type Auth struct {
	Id              int    `form:"id"  json:"id"`
	Name            string `form:"name" json:"name" binding:"required"`
	Telephone       string `form:"telephone" json:"telephone" binding:"required"`
	Group           string `form:"group" json:"group" binding:"required"`
	Username        string `form:"username"  json:"username" binding:"required"`
	Password        string `form:"password"  json:"password" binding:"required"`
	IsAdministrator int    `form:"is_administrator" json:"is_administrator"`
	CreatedBy       string `form:"created_by"  json:"created_by"`
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

func (a *Auth) GetNameByUsername() (string, string, error) {
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

func (a *Auth) IsExistUser() (bool, error) {
	return models.IsExistUser(a.Name, a.Username)
}
