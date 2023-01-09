package errmsg

const (
	Success = 200
	Error   = 500

	// ErrorCodeSms code= 1000... 用户模块的错误
	ErrorCodeSms = 1001

	ErrorPasswordWrong = 1002
	ErrorUserNotExist  = 1003
	ErrorUserNoRight   = 1004
	ErrorUserNoEmail   = 1005
	ErrorUserYesExist  = 1006

	ErrorTokenExist     = 1007
	ErrorTokenRuntime   = 1008
	ErrorTokenWrong     = 1009
	ErrorTokenTypeWrong = 1010

	// ErrorCatenameUsed code= 3000... 统计模块的错误
	ErrorCatenameUsed = 3001
	ErrorCateNotExist = 3002
)

var codeMsg = map[int]string{
	Success: "OK",
	Error:   "请求失败",

	ErrorCodeSms: "验证码错误",

	ErrorPasswordWrong: "密码错误",
	ErrorUserNotExist:  "用户不存在",
	ErrorUserNoRight:   "该用户无权限",
	ErrorUserNoEmail:   "邮箱错误",
	ErrorUserYesExist:  "用户已存在",

	ErrorTokenExist:     "TOKEN不存在,请重新登陆",
	ErrorTokenRuntime:   "TOKEN已过期,请重新登陆",
	ErrorTokenWrong:     "TOKEN不正确,请重新登陆",
	ErrorTokenTypeWrong: "TOKEN格式错误,请重新登陆",

	ErrorCatenameUsed: "该分类已存在",
	ErrorCateNotExist: "该分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
