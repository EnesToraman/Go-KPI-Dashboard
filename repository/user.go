package repository

import (
	"database/sql"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

type UserRepository struct{}

// Login creates a token if the credentials given are true.
func (lr *UserRepository) GetUserEmailRole(db *sql.DB, userCredentialsDTO model.UserCredentialsDTO) (model.UserEmailRoleDTO, error) {
	query := `
	SELECT hashed_password, role FROM user
	WHERE email = ?
	`
	row := db.QueryRow(query, userCredentialsDTO.Email)
	var hashedPassword string
	var role string
	err := row.Scan(&hashedPassword, &role)
	if err != nil {
		return model.UserEmailRoleDTO{}, err
	}
	err = utils.CheckPassword(userCredentialsDTO.Password, hashedPassword)
	if err != nil {
		return model.UserEmailRoleDTO{}, err
	}
	userEmailRoleDTO := model.UserEmailRoleDTO{
		Email: userCredentialsDTO.Email,
		Role:  role,
	}
	return userEmailRoleDTO, nil
}

// SignUp signs the user up with email and password.
func (ur *UserRepository) CreateUser(db *sql.DB, userCredentialsDTO model.UserCredentialsDTO) (sql.Result, error) {
	query := `
	INSERT INTO user (email, hashed_password, role) 
	VALUES (?, ?, 'STAFF')
	`
	hashedPassword, err := utils.HashPassword(userCredentialsDTO.Password)
	if err != nil {
		return nil, err
	}

	res, err := db.Exec(query, userCredentialsDTO.Email, hashedPassword)
	if err != nil {
		return nil, err
	}
	return res, nil
}
