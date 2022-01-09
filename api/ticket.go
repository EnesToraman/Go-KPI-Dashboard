package api

import (
	"net/http"

	"github.com/EnesToraman/Go-KPI-Dashboard/config"
	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
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
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	for rows.Next() {
		if err := rows.Scan(&ticketdata.NumberOfTickets, &ticketdata.Class, &ticketdata.DepPlace, &ticketdata.Date); err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		ticketdatas = append(ticketdatas, ticketdata)
	}
	if err := rows.Err(); err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	utils.RespondWithSuccess(ticketdatas, w)
}

// GetRevenue retrieves the relevant data with join operations, then scans and returns.
func GetRevenue(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	var ticketDTO model.TicketDTO
	var ticketDTOs []model.TicketDTO
	query := `
	SELECT SUM(t.price), DATE(f.dep_time) AS date FROM ticket AS t
	INNER JOIN flight AS f ON t.flight_id = f.flight_id
	GROUP BY date
	ORDER BY date
	`
	db := config.DB()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	for rows.Next() {
		if err := rows.Scan(&ticketDTO.Revenue, &ticketDTO.Date); err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		ticketDTOs = append(ticketDTOs, ticketDTO)
	}
	if err := rows.Err(); err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)

	}
	utils.RespondWithSuccess(ticketDTOs, w)
}

// GetAverageTicketPrice retrieves the relevant data with join operations, then scans and returns.
func GetAverageTicketPrice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	var ticketDTO model.TicketDTO
	var ticketDTOs []model.TicketDTO
	query := `
	SELECT DATE(f.dep_time) AS date, ROUND(AVG(t.price),0) FROM ticket AS t
	INNER JOIN flight AS f ON t.flight_id = f.flight_id
	GROUP BY date
	ORDER BY date
	`
	db := config.DB()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	for rows.Next() {
		if err := rows.Scan(&ticketDTO.Date, &ticketDTO.Revenue); err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)

		}
		ticketDTOs = append(ticketDTOs, ticketDTO)
	}
	if err := rows.Err(); err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)

	}
	utils.RespondWithSuccess(ticketDTOs, w)
}

// GetTicketClass retrieves the relevant data with join operations, then scans and returns.
func GetTicketClass(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	var ticketClassDTO model.TicketClassDTO
	var ticketClassDTOs []model.TicketClassDTO
	query := `
	SELECT class, COUNT(ticket_id) FROM ticket
	GROUP BY class
	`
	db := config.DB()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	for rows.Next() {
		if err := rows.Scan(&ticketClassDTO.Class, &ticketClassDTO.NumberOfTickets); err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		ticketClassDTOs = append(ticketClassDTOs, ticketClassDTO)
	}
	if err := rows.Err(); err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	utils.RespondWithSuccess(ticketClassDTOs, w)
}
