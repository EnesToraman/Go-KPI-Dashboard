package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/EnesToraman/Go-KPI-Dashboard/config"
	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

// SignUp signs the user up with email and password.
func SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pwd := user.Password
	email := user.Email
	hashedPassword, err := utils.HashPassword(pwd)
	if err != nil {
		fmt.Printf("CreateUser: %v", err)
	}

	db := config.DB()
	_, err2 := db.Exec("INSERT INTO user (email, hashed_password, role) VALUES (?, ?, 'STAFF')", email, hashedPassword)
	if err2 != nil {
		fmt.Printf("CreateUser: %v", err2)
	}
}
