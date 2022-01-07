package model

type Ticket struct {
	Ticket_id    int    `json:"ticketId"`
	Passenger_id int    `json:"passengerId"`
	Flight_id    int    `json:"flightId"`
	Seat         string `json:"seat"`
	Class        string `json:"class"`
	Price        int    `json:"price"`
}

type TicketDataModel struct {
	NumberOfPassengers int    `json:"numberOfPassengers"`
	TotalPrice         int    `json:"totalPrice"`
	Class              string `json:"class"`
	DepPlace           string `json:"depPlace"`
	Date               string `json:"date"`
}
