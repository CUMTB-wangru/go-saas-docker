package user

// SmsPhone 短信
type SmsPhone struct {
	Phone string `json:"phone" binding:"required"`
}

// Register 注册
type Register struct {
	UserName     string `json:"name"`
	Password     string `json:"password"`
	Phone        string `json:"phone" binding:"required"`
	Email        string `json:"email"`
	Code         string `json:"code"`
	SurePassword string `json:"sure_password"`
}

// Login 密码登录
type Login struct {
	UserName string `json:"name"`
	Password string `json:"password"`
}
