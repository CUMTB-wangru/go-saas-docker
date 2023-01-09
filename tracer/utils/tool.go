package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
	"time"
)

// NumberRandom 随机验证码
func NumberRandom() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// VerifyMobileFormat 电话校验: 目前用不上
func VerifyMobileFormat(mobileNum string) bool {
	regular := `^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// VerifyEmailFormat 邮箱校验
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)@\w+([-.]\w+).\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// EncryptPsaaword 密码加密
func EncryptPsaaword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		panic(err)
	}
	pwd := string(hash)
	return pwd
}

// ComparePassword 验证密码
func ComparePassword(sqlPassword, password string) bool {
	sqlPwd := []byte(sqlPassword)
	pwd := []byte(password)
	err := bcrypt.CompareHashAndPassword(sqlPwd, pwd)
	if err != nil {
		return false
	}
	return true
}
