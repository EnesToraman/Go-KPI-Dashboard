package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EnesToraman/Go-KPI-Dashboard/config"
	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/julienschmidt/httprouter"
)

func GetTicketData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	var passenger model.TicketDataModel
	var passengers []model.TicketDataModel
	query := `
	SELECT COUNT(t.passenger_id), SUM(t.price), t.class, f.dep_place, DATE(f.dep_time) AS date FROM ticket AS t
	INNER JOIN flight AS f ON t.flight_id = f.flight_id
    GROUP BY date, dep_place, class
    ORDER BY date
	`
	db := config.DB()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("TicketDataModel: %v", err)
	}
	for rows.Next() {
		if err := rows.Scan(&passenger.NumberOfPassengers, &passenger.TotalPrice, &passenger.Class, &passenger.DepPlace, &passenger.Date); err != nil {
			fmt.Printf("TicketDataModel: %v", err)
		}
		passengers = append(passengers, passenger)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("TicketDataModel: %v", err)
	}
	json.NewEncoder(w).Encode(passengers)
}
