package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/EnesToraman/Go-KPI-Dashboard/controller"
	"github.com/EnesToraman/Go-KPI-Dashboard/driver"
)

func New() *httprouter.Router {
	r := httprouter.New()
	// Configuration for pre-flight requests.
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
	InitRoutes(r)
	return r
}

func InitRoutes(r *httprouter.Router) {
	db := driver.DB()
	ec := controller.EmployeeController{}
	tc := controller.TicketController{}
	pc := controller.PlaneController{}
	uc := controller.UserController{}

	// Endpoints for authentication and authorization process
	r.POST("/login", uc.Login(db))
	r.GET("/authUser", controller.AuthUser)
	r.POST("/logout", controller.Logout)

	r.POST("/signUp", uc.CreateUser(db))
	r.GET("/ticketCountGroupByDate", tc.GetTicketCountGroupByDate(db))
	r.GET("/totalPriceGroupByDate", tc.GetTotalPriceGroupByDate(db))
	r.GET("/averageTicketPriceGroupByDate", tc.GetTicketAveragePriceGroupByDate(db))
	r.GET("/ticketCountGroupByClass", tc.GetTicketCountGroupByClass(db))
	r.GET("/employeeCountGroupByTitle", ec.GetEmployeeCountGroupByTitle(db))
	r.GET("/employeeAverageSalaryGroupByTitle", ec.GetEmployeeAverageSalaryGroupByTitle(db))
	r.GET("/planeCountGroupByAirline", pc.GetPlaneCountGroupByAirline(db))
}
