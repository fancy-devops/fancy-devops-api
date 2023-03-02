package requests

type HookReq struct {
	EventName  string         `json:"event_name"`
	Ref        string         `json:"ref"`
	UserId     int64          `json:"user_id"`
	UserName   string         `json:"user_name"`
	UserEmail  string         `json:"user_email"`
	Repository HookRepository `json:"repository"`
}

type HookRepository struct {
	Name    string `json:"name"`
	HttpUrl string `json:"git_http_url"`
	SshUrl  string `json:"git_ssh_url"`
}
