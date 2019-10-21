package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email       string `grom:"column:email;"` // 邮箱
	MobilePhone string                        // 电话
	PassWord    string                        // 密码
}
