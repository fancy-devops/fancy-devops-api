package db

const (
	Role_Guest = 0
	Role_User  = 1
	Role_Admin = 2
)

type User struct {
	BaseModel

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
	Role     int    `json:"role"`
}
