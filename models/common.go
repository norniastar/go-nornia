package models

import (
	. "go-nornia/utils/db/mysql"
)

var tablePrefix = "sl"

func InitModels() {
	_ = DB.AutoMigrate(&User{})
	user := &User{
		Uid:      "111",
		Tel:      "123456",
		Name:     "name",
		Password: "1",
		Status:   1,
	}
	var count int64
	_ = DB.Table(user.TableName()).Where("tel = ?", user.Tel).Count(&count)
	if count == 0 {
		DB.Create(user)
	}
}
