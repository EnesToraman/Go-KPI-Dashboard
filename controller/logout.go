package controller

import (
	"net/http"
	"time"

	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
	"github.com/julienschmidt/httprouter"
)

// Logout sets the token value to empty.
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.CORS(w, r)
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
}
