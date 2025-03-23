package models

import "time"

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
type SeatAssignments struct {
	PassengerName string `json:"passenger_name"`
	PassengerType string `json:"passenger_type"` // adult,infant,elderly etc
	SeatNumber    string `json:"seat_number"`
	SeatType      string `json:"seat_type"` //window,aisle,middle
}
type BookedTicket struct {
	User                User              `json:"user"`
	Airline             Airline           `json:"airline"`
	NumberOfTicket      TicketTypes       `json:"numberofTickets"`
	TicketClass         TicketClass       `json:"ticket_class"`
	BookingStatus       string            `json:"Booking_status"`
	BookingDate         time.Time         `json:"booking_date"`
	TotalPrice          float64           `json:"total_price"`
	PaymentStatus       string            `json:"payment_status"`
	CheckInStatus       bool              `json:"check_in_status"`
	SeatAssignments     []SeatAssignments `json:"seat_assignments"`
	SpecialRequests     []string          `json:"special_requests"`
	CancellationPolicay bool              `json:"cancellation_policy"`
	ChangeFee           float64           `json:"change_fee"`
}
