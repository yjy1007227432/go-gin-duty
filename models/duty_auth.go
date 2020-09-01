package models

import (
	"time"
)

type DutyAuth struct {
	Id              int       `xorm:"not null pk autoincr INT(10)"`
	Name            string    `xorm:"default '' comment('姓名') VARCHAR(50)"`
	Telephone       string    `xorm:"default '' comment('电话') VARCHAR(50)"`
	Group           string    `xorm:"default '' comment('所属组：计费:calculate  crm:crm') VARCHAR(50)"`
	Username        string    `xorm:"default '' comment('账号') VARCHAR(50)"`
	Password        string    `xorm:"default '' comment('密码') VARCHAR(50)"`
	IsAdministrator int       `xorm:"default 0 comment('是否管理员，0：否，1：是') TINYINT(3)"`
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
		IsAdministrator: data["is_administrator"].(int),
		CreatedOn:       time.Now(),
		CreatedBy:       data["created_by"].(string),
	}
	if err := db.Create(&auth).Error; err != nil {
		return err
	}

	return nil
}
