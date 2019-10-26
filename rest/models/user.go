package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email       string `grom:"column:email;"` // 邮箱
	MobilePhone string                        // 电话
	PassWord    string                        // 密码
}

func (u *User) CreateByEmail() error {
	user := &User{Email: u.Email, PassWord: u.PassWord}
	return GetDB().Create(&user).Error
}

func (u *User) CreateByPhone() error {
	user := &User{Email: u.Email, MobilePhone: u.MobilePhone}
	return GetDB().Create(&user).Error
}

func (u *User) Add() error {
	return GetDB().Create(u).Error
}

func (u *User) FindById() error {
	return GetDB().Find(u, u.Model.ID).Error
}

func (u *User) FindByEmail() error {
	return GetDB().Find(u, &User{Email: u.Email}).Error
}
