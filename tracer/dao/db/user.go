package db

import (
	. "tracer/model"
)

// SetUserInfo 创建用户
func SetUserInfo(userName, password, phone, email string) {
	userInfo := UserInfo{
		UserName: userName,
		Password: password,
		Phone:    phone,
		Email:    email,
	}
	DB.Create(&userInfo)
}

// CheckUser 查询用户
func CheckUser(phone string) int64 {
	var userInfo UserInfo
	var count int64
	DB.Where("phone = ?", phone).Find(&userInfo).Count(&count)
	return count
}

// CheckLogin 登录查询
func CheckLogin(msg string) (UserInfo, int64) {
	var userInfo UserInfo
	var count int64
	DB.Where("phone = ?", msg).Or("email = ?", msg).Find(&userInfo).Count(&count)
	return userInfo, count
}
