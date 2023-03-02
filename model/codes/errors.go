package codes

const (
	Common_Success      = 200
	Common_BadRequest   = 400
	Common_Unauthorized = 401
	Common_Forbidden    = 403
	Common_Error        = 500

	Error_Failed = 10000

	Error_User_Exist         = 10001
	Error_User_NotExist      = 10002
	Error_User_WrongPassword = 10003
	Error_User_WrongCode     = 10004
	Error_User_Guest         = 10005

	Error_User_AuthTokenNotExist   = 10051
	Error_User_AuthTokenParseError = 10052
	Error_User_AuthTokenExpired    = 10053
	Error_User_AuthTokenGenFailed  = 10054

	Error_Dir_NotExist = 20001
	Error_Dir_Exist    = 20002

	Error_File_NotExist = 30001
	Error_File_Exist    = 30002
)

var MsgFlags = map[int]string{
	Common_Success:      "ok",
	Common_BadRequest:   "错误的请求",
	Common_Unauthorized: "未登录",
	Common_Forbidden:    "无权限",
	Common_Error:        "系统异常",

	Error_Failed: "操作失败",

	Error_User_Exist:         "用户已存在",
	Error_User_NotExist:      "用户不存在",
	Error_User_WrongPassword: "密码错误",
	Error_User_WrongCode:     "code错误",
	Error_User_Guest:         "暂无法登录，请联系管理员进行审核",

	Error_User_AuthTokenNotExist:   "认证token不存在",
	Error_User_AuthTokenParseError: "认证token解析失败",
	Error_User_AuthTokenExpired:    "认证token已过期",
	Error_User_AuthTokenGenFailed:  "认证token生成失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Common_Error]
}
