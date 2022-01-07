package model

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"Password"`
	Role     string `json:"role"`
}
