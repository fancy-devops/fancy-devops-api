package db

type Operation_Log struct {
	BaseModel

	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}
