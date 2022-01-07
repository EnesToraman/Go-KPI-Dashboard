package router

import (
	"net/http"

	"github.com/EnesToraman/Go-KPI-Dashboard/api"
	"github.com/EnesToraman/Go-KPI-Dashboard/api/auth"
	"github.com/julienschmidt/httprouter"
)

func New() *httprouter.Router {
	r := httprouter.New()
	// Configure for pre-flight requests.
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Content-Type", "application/json")
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
			header.Set("Access-Control-Allow-Credentials", "true")
			header.Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	Routes(r)
	return r
}

func Routes(r *httprouter.Router) {
	// Endpoints for authentication and authorization process
	r.POST("/login", auth.Login)
	r.GET("/authUser", auth.AuthUser)
	r.POST("/logout", auth.Logout)

	r.POST("/sign-up", api.SignUp)
	r.GET("/ticketData", api.GetTicketData)

}
