package main

import (
	"log"
	"net/http"

	"github.com/EnesToraman/Go-KPI-Dashboard/router"
)

func main() {

	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
