package errmsg

const (
	SUCCSE = 200
	ERROR  = 500
	//code=1000 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	//code=2000 文章模块的错误
	ERROR_CATENAME_USED = 2001
	//code=300 分类模块的错误
	ERROR_ART_NOT_EXIST = 3001
)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "token不存在",
	ERROR_TOKEN_RUNTIME:    "token过期",
	ERROR_TOKEN_WRONG:      "token错误",
	ERROR_TOKEN_TYPE_WRONG: "token格式不正确",
	ERROR_CATENAME_USED:    "该分类已经存在",
	ERROR_ART_NOT_EXIST:    "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
