package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type DutyAuth struct {
	Id              int       `form:"id"  json:"id"`
	Name            string    `form:"name" json:"name"`
	Telephone       string    `form:"telephone" json:"telephone"`
	Group           string    `form:"group" json:"group"`
	Username        string    `form:"username"  json:"username"`
	Password        string    `form:"password"  json:"password"`
	IsAdministrator int       `form:"is_administrator" json:"is_administrator"`
	CreatedOn       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	CreatedBy       string    `xorm:"default '' comment('创建人') VARCHAR(50)"`
	ModifiedOn      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	ModifiedBy      string    `xorm:"default '' comment('修改人') VARCHAR(50)"`
	Backup1         string    `xorm:"default '' VARCHAR(50)"`
	Backup2         string    `xorm:"default '' VARCHAR(50)"`
}

func AddAuth(data map[string]interface{}) error {
	auth := DutyAuth{
		Name:            data["name"].(string),
		Telephone:       data["telephone"].(string),
		Group:           data["group"].(string),
		Username:        data["username"].(string),
		Password:        data["password"].(string),
		IsAdministrator: 0,
		CreatedOn:       time.Now(),
		CreatedBy:       data["created_by"].(string),
	}
	err := db.Create(&auth).Error
	if err != nil {
		log.Printf("db.Create err: %v", err)
		return err
	}
	return nil
}

//绑定表名
//func (v DutyAuth) TableName() string {
//	return "duty_auth"
//}

func UpdateAuth(data map[string]interface{}) error {
	auth := DutyAuth{}
	err := db.Model(&auth).Updates(data)
	if err != nil {
		log.Printf("UpdateAuth err: %v", err)
	}

	return nil
}

func IsAdmin(username string) (int, error) {
	auth := DutyAuth{
		Username: username,
	}

	err := db.Select("is_administrator").Where(&auth).Find(&auth).Error

	if err != nil {
		log.Printf("IsAdmin err: %v", err)
	}

	return auth.IsAdministrator, nil
}

func GetGroup(name string) (string, error) {
	auth := DutyAuth{
		Name: name,
	}
	err := db.Select("`group`").Where(&auth).Find(&auth).Error

	if err != nil {
		log.Printf("GetGroup err: %v", err)
	}

	return auth.Group, nil
}

func DeleteAuthById(id int) error {
	auth := DutyAuth{Id: id}
	err := db.Delete(&auth)
	if err != nil {
		log.Printf("DeleteAuthById err: %v", err)
	}
	return nil
}

func DeleteAuthByUsername(username string) error {
	err := db.Where("username = ?", username).Delete(&DutyAuth{})
	if err != nil {
		log.Printf("DeleteAuthByUsername err: %v", err)
	}
	return nil
}

func CheckAuth(username, password string) (bool, error) {
	var auth DutyAuth
	err := db.Select("id").Where(DutyAuth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.Id > 0 {
		return true, nil
	}

	return false, nil
}

func GetNameByUsername(username string) (string, string, error) {
	var auth DutyAuth
	err := db.Select("id,name,telephone,`group`").Where(&DutyAuth{Username: username}).First(&auth).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return "", "", err
	}

	if auth.Id > 0 {
		return auth.Name, auth.Group, nil
	}

	return "", "", nil
}

func IsExistUser(name, userName string) (bool, error) {

	var auth DutyAuth
	err := db.Select("id").Where(DutyAuth{Name: name}).Or(DutyAuth{Username: userName}).First(&auth).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.Id > 0 {
		return true, nil
	}

	return false, nil
}

func IsExistName(username string, password string) (bool, error) {

	var auth DutyAuth
	err := db.Select("id").Where(DutyAuth{Username: username, Password: password}).First(&auth).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.Id > 0 {
		return false, nil
	}

	return true, nil
}
