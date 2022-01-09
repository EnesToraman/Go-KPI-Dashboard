package model

type Plane struct {
	ID          int    `json:"id"`
	AirlineName string `json:"airlineName"`
	Code        string `json:"code"`
	Capacity    int    `json:"capacity"`
}

type PlaneCountGroupByAirlineDTO struct {
	AirlineName    string `json:"airlineName"`
	NumberOfPlanes int    `json:"numberOfPlanes"`
}
