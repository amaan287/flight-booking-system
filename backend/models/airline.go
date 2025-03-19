package models

import "time"

type Class struct {
	Economy        int `json:"economy"`
	PremiumEconomy int `json:"premium_economy"`
	Businessclass  int `json:"business_class"`
	FirstClass     int `json:"first_class"`
}
type Airline struct {
	Id               int       `json:"id"`
	From             string    `json:"from"`
	To               string    `json:"to"`
	Date             time.Time `json:"date"`
	AirportName      string    `json:"airport_name"`
	NoOfSeatsInClass Class     `json:"noOfSeatInclass"`
	Name             string    `json:"name"`
}
