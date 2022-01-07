package auth

import (
	"net/http"
	"time"

	"github.com/EnesToraman/Go-KPI-Dashboard/config"
	"github.com/julienschmidt/httprouter"
)

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
}
