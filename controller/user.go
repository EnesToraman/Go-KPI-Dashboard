package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/repository"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func (uc *UserController) Login(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		utils.CORS(w, r)
		var UserCredentialsDTO model.UserCredentialsDTO
		userRepo := repository.UserRepository{}

		err := json.NewDecoder(r.Body).Decode(&UserCredentialsDTO)
		if err != nil {
			utils.RespondWithError(err, http.StatusBadRequest, w)
		}

		userEmailRoleDTO, err := userRepo.GetUserEmailRole(db, UserCredentialsDTO)
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}

		expirationTime := time.Now().Add(time.Hour * 24)
		claims := &model.Claims{
			Email: userEmailRoleDTO.Email,
			Role:  userEmailRoleDTO.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
		})
	}
}

func (uc *UserController) CreateUser(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		utils.CORS(w, r)
		var userCredentialsDTO model.UserCredentialsDTO
		err := json.NewDecoder(r.Body).Decode(&userCredentialsDTO)
		if err != nil {
			utils.RespondWithError(err, http.StatusBadRequest, w)
		}
		userRepo := repository.UserRepository{}
		res, err := userRepo.CreateUser(db, userCredentialsDTO)
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		utils.RespondWithSuccess(res, w)
	}
}
