package repository

import (
	"database/sql"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
)

type TicketRepository struct{}

func (tr *TicketRepository) GetTicketCountGroupByDate(db *sql.DB) ([]model.TicketCountGroupByDateDTO, error) {
	var ticketdata model.TicketCountGroupByDateDTO
	var ticketdatas []model.TicketCountGroupByDateDTO
	query := `
	SELECT COUNT(t.ticket_id), t.class, f.dep_place, DATE(f.dep_time) AS date FROM ticket AS t
	INNER JOIN flight AS f ON t.flight_id = f.flight_id
    GROUP BY date, dep_place, class
    ORDER BY date
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&ticketdata.NumberOfTickets, &ticketdata.Class, &ticketdata.DepPlace, &ticketdata.Date); err != nil {
			return nil, err
		}
		ticketdatas = append(ticketdatas, ticketdata)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ticketdatas, nil
}

func (tr *TicketRepository) GetTotalPriceGroupByDate(db *sql.DB) ([]model.TicketTotalPriceGroupByDateDTO, error) {
	var ticketDTO model.TicketTotalPriceGroupByDateDTO
	var ticketDTOs []model.TicketTotalPriceGroupByDateDTO
	query := `
	SELECT SUM(t.price), DATE(f.dep_time) AS date FROM ticket AS t
	INNER JOIN flight AS f ON t.flight_id = f.flight_id
	GROUP BY date
	ORDER BY date
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&ticketDTO.Revenue, &ticketDTO.Date); err != nil {
			return nil, err
		}
		ticketDTOs = append(ticketDTOs, ticketDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ticketDTOs, nil
}

func (tr *TicketRepository) GetTicketAveragePrice(db *sql.DB) ([]model.TicketAveragePriceGroupByDateDTO, error) {
	var ticketDTO model.TicketAveragePriceGroupByDateDTO
	var ticketDTOs []model.TicketAveragePriceGroupByDateDTO
	query := `
	SELECT DATE(f.dep_time) AS date, ROUND(AVG(t.price),0) FROM ticket AS t
	INNER JOIN flight AS f ON t.flight_id = f.flight_id
	GROUP BY date
	ORDER BY date
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&ticketDTO.Date, &ticketDTO.AvgPrice); err != nil {
			return nil, err
		}
		ticketDTOs = append(ticketDTOs, ticketDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ticketDTOs, nil
}

func (tr *TicketRepository) GetTicketCountGroupByClass(db *sql.DB) ([]model.TicketCountGroupByClassDTO, error) {
	var ticketClassDTO model.TicketCountGroupByClassDTO
	var ticketClassDTOs []model.TicketCountGroupByClassDTO
	query := `
	SELECT class, COUNT(ticket_id) FROM ticket
	GROUP BY class
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&ticketClassDTO.Class, &ticketClassDTO.NumberOfTickets); err != nil {
			return nil, err
		}
		ticketClassDTOs = append(ticketClassDTOs, ticketClassDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ticketClassDTOs, nil
}
