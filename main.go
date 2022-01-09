package main

import (
	"net/http"

	"github.com/EnesToraman/Go-KPI-Dashboard/router"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

func main() {
	router := router.New()
	err := http.ListenAndServe(":8080", router)
	utils.CheckError(err)
}
