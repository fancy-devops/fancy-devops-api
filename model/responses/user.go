package responses

type UserListRes struct {
	BaseRes
	Data PagedUserList `json:"data"`
}

type UserDetailRes struct {
	BaseRes
	Data UserDetailModel `json:"data"`
}

type UserLoginRes struct {
	BaseRes
	Data UserLoginModel `json:"data"`
}

type UserDetailModel struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  int    `json:"role"`
}

type PagedUserList struct {
	Total int64             `json:"total"`
	List  []UserDetailModel `json:"list"`
}

type UserLoginModel struct {
	Token string `json:"token"`
}
