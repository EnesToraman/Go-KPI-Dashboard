package main

import (
	"net/http"

	"github.com/EnesToraman/Go-KPI-Dashboard/router"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

func main() {

	r := router.New()
	err := http.ListenAndServe(":8080", r)
	utils.CheckError(err)
}
