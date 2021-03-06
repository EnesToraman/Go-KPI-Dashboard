package controller

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

// AuthUser checks whether request token is valid or not, if so, returns claims.
func AuthUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.CORS(w, r)
	cookie, err := r.Cookie("token")
	if err != nil {
		utils.RespondWithError(err, http.StatusBadRequest, w)
	}
	tokenStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTKEY")), nil
	})
	if err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	utils.RespondWithSuccess(claims, w)
}
