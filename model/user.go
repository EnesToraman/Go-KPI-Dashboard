package model

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserCredentialsDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEmailRoleDTO struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}
