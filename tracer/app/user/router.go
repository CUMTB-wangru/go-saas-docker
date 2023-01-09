package user

import (
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.POST("/api/user/sms", sendSms)
	r.POST("/api/user/register", registerHandler)
	r.POST("/api/user/login", loginHandler)
	r.POST("/api/user/login_sms", loginSmsHandler)
}
