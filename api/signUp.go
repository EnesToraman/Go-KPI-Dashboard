package api

import (
	"encoding/json"
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
	query := `
	INSERT INTO user (email, hashed_password, role) 
	VALUES (?, ?, 'STAFF')
	`
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithError(err, http.StatusBadRequest, w)
	}
	pwd := user.Password
	email := user.Email
	hashedPassword, err := utils.HashPassword(pwd)
	if err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}

	db := config.DB()
	res, err := db.Exec(query, email, hashedPassword)
	if err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	utils.RespondWithSuccess(res, w)
}
