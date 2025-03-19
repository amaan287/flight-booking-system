package models

type TicketTypes struct {
	Adult    int `json:"adult"`
	Childern int `json:"children"`
	Infant   int `json:"infant"`
	Elderly  int `json:"elderly"`
}
type TicketClass struct {
	Economy        bool `json:"economy"`
	PremiumEconomy bool `json:"premium_economy"`
	Businessclass  bool `json:"business_class"`
	FirstClass     bool `json:"first_class"`
}
type BookedTicket struct {
	User           User        `json:"user"`
	Airline        Airline     `json:"airline"`
	NumberOfTicket TicketTypes `json:"numberofTickets"`
	TicketClass
}
