package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EnesToraman/Go-KPI-Dashboard/config"
	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/julienschmidt/httprouter"
)

// GetTicketData retrieves the relevant data with join operations, then scans and returns.
func GetTicketData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	var ticketdata model.TicketDataModel
	var ticketdatas []model.TicketDataModel
	query := `
	SELECT COUNT(t.ticket_id), t.class, f.dep_place, DATE(f.dep_time) AS date FROM ticket AS t
	INNER JOIN flight AS f ON t.flight_id = f.flight_id
    GROUP BY date, dep_place, class
    ORDER BY date
	`
	db := config.DB()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("TicketDataModel: %v", err)
		return
	}
	for rows.Next() {
		if err := rows.Scan(&ticketdata.NumberOfTickets, &ticketdata.Class, &ticketdata.DepPlace, &ticketdata.Date); err != nil {
			fmt.Printf("TicketDataModel: %v", err)
			return
		}
		ticketdatas = append(ticketdatas, ticketdata)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("TicketDataModel: %v", err)
		return
	}
	json.NewEncoder(w).Encode(ticketdatas)
}
