package models

import (
	"go-nornia/utils/db/mysql"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Uid       string `gorm:"column:uid;not null;type:int;comment:用户ID" json:"uid"`
	Tel       string `gorm:"column:tel;not null;type:varchar(20);comment:手机号" json:"tel"`
	Name      string `gorm:"column:name;not null;type:varchar(20);comment:用户名" json:"name"`
	Password  string `gorm:"column:password;not null;type:varchar(20);comment:密码" json:"password"`
	Status    int    `gorm:"column:status;not null;type:tinyint;comment:'状态 0-禁用 1-可用'" json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *User) TableName() string {
	return tablePrefix + "_user"
}

func NewUserModel() *User {
	return &User{}
}

func (m *User) GetByTel(tel string) error {
	db := mysql.DB

	m.Tel = tel
	err := db.First(m, "tel")
	return err.Error
}
