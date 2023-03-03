package requests

type GetUserListReq struct {
	Role      int `json:"role"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type SendVerifyEmailReq struct {
	Email string `json:"email"`
}

type UserRegisterReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
	Code  string `json:"code"`
}

type UserLoginReq struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type UserResetPwdReq struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type UserChangePwdReq struct {
	OldPwd string `json:"oldpwd"`
	NewPwd string `json:"newpwd"`
}
