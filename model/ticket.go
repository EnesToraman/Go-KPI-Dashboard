package model

type Ticket struct {
	Ticket_id    int    `json:"ticketId"`
	Passenger_id int    `json:"passengerId"`
	Flight_id    int    `json:"flightId"`
	Seat         string `json:"seat"`
	Class        string `json:"class"`
	Price        int    `json:"price"`
}

type TicketCountGroupByDateDTO struct {
	NumberOfTickets int    `json:"numberOfPassengers"`
	Class           string `json:"class"`
	DepPlace        string `json:"depPlace"`
	Date            string `json:"date"`
}

type TicketTotalPriceGroupByDateDTO struct {
	Date    string `json:"date"`
	Revenue int    `json:"totalPrice"`
}

type TicketAveragePriceGroupByDateDTO struct {
	Date     string `json:"date"`
	AvgPrice string `json:"avgPrice"`
}

type TicketCountGroupByClassDTO struct {
	Class           string `json:"class"`
	NumberOfTickets int    `json:"numberOfTickets"`
}
