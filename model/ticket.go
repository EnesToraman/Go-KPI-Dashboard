package model

type TicketDataModel struct {
	NumberOfPassengers int    `json:"numberOfPassengers"`
	TotalPrice         int    `json:"totalPrice"`
	Class              string `json:"class"`
	DepPlace           string `json:"depPlace"`
	Date               string `json:"date"`
}
