package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"

	"github.com/EnesToraman/Go-KPI-Dashboard/config"
	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

// Login creates a token if the credentials given are true.
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	config.CORS(w, r)
	var credentials model.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	email := credentials.Email
	password := credentials.Password

	db := config.DB()
	defer db.Close()
	row := db.QueryRow("SELECT hashed_password, role FROM user WHERE email = ?", email)
	var hashedPassword string
	var role string
	err = row.Scan(&hashedPassword, &role)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = utils.CheckPassword(password, hashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 60)

	claims := &model.Claims{
		Email: credentials.Email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	})
}
