package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tracer/dao/db"
	"tracer/dao/redis"
	"tracer/utils"
	"tracer/utils/errmsg"
	"tracer/utils/sms"
)

//  短信
func sendSms(c *gin.Context) {
	var smsPhone SmsPhone
	if err := c.ShouldBind(&smsPhone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.Error,
			"data":    "",
			"message": errmsg.GetErrMsg(errmsg.Error),
		})
		return
	}
	// 生成随机6位验证码
	num := utils.NumberRandom()
	// 存储到redis中，有效期60秒，加上下标register_与登录作区分
	rdbNum := redis.RdbStore{}.Ttl("register_", smsPhone.Phone)
	if rdbNum >= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ErrorCodeSms,
			"data":    "",
			"message": "短信已发送，请" + fmt.Sprintf("%f", rdbNum) + "秒以后再点击发送",
		})
		return
	}
	// 短信验证码发送
	sms.Sms(smsPhone.Phone, num)
	// redis存储
	redis.RdbStore{}.Set("register_", smsPhone.Phone, num)

	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.Success,
		"data":    "",
		"message": "发送成功，请查看手机短信",
	})
}

// 注册
func registerHandler(c *gin.Context) {
	var register Register
	if err := c.ShouldBind(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.Error,
			"data":    "",
			"message": errmsg.GetErrMsg(errmsg.Error),
		})
		return
	}

	// 检查用户是否已注册
	count := db.CheckUser(register.Phone)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ErrorUserYesExist,
			"data":    "",
			"message": errmsg.GetErrMsg(errmsg.ErrorUserYesExist),
		})
		return
	}

	// 判断短信验证码（手机号）
	num := redis.RdbStore{}.Get("register_", register.Phone)
	if num == "" || num != register.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ErrorCodeSms,
			"data":    "",
			"message": errmsg.GetErrMsg(errmsg.ErrorCodeSms),
		})
		return
	}

	// 密码校验
	if register.Password != register.SurePassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ErrorPasswordWrong,
			"data":    "",
			"message": errmsg.GetErrMsg(errmsg.ErrorPasswordWrong),
		})
		return
	}
	// 邮箱校验
	check := utils.VerifyEmailFormat(register.Email)
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ErrorUserNoEmail,
			"data":    "",
			"message": errmsg.GetErrMsg(errmsg.ErrorUserNoEmail),
		})
		return
	}
	// 密码加盐加密
	password := utils.EncryptPsaaword(register.Password)
	// 写入数据库
	db.SetUserInfo(
		register.UserName, password, register.Phone, register.Email,
	)

	// jwt签发
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.Success,
		"data":    "",
		"message": "用户注册成功，已跳转首页",
	})
}

// 密码登录
func loginHandler(c *gin.Context) {
	var login Login
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.Error,
			"data":    "",
			"message": errmsg.GetErrMsg(errmsg.Error),
		})
		return
	}

	userInfo, count := db.CheckLogin(login.UserName)

	// 校验用户是否存在
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ErrorUserNotExist,
			"data":    "",
			"message": errmsg.GetErrMsg(errmsg.ErrorUserNotExist),
		})
		return
	}
	// 密码校验
	flag := utils.ComparePassword(userInfo.Password, login.Password)
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ErrorPasswordWrong,
			"data":    "",
			"message": errmsg.GetErrMsg(errmsg.ErrorPasswordWrong),
		})
		return
	}
	// 腾讯防水墙
	// jwt签发
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"data":    "",
		"message": "登陆成功",
	})
}

// 短信登录
func loginSmsHandler(c *gin.Context) {
	return
}
