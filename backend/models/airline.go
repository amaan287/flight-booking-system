package models

import "time"

type Class struct {
	Economy        int `json:"economy"`
	PremiumEconomy int `json:"premium_economy"`
	Businessclass  int `json:"business_class"`
	FirstClass     int `json:"first_class"`
}
type ClassPrices struct {
	Economy        float64 `json:"economy"`
	PremiumEconomy float64 `json:"premium_economy"`
	Businessclass  float64 `json:"business_class"`
	FirstClass     float64 `json:"first_class"`
}
type FlightStatus string

// ^ shift this to the code where you gonna use flightstatus
const (
	Scheduled FlightStatus = "scheduled"
	Boarding  FlightStatus = "boarding"
	Departed  FlightStatus = "departed"
	Delay     FlightStatus = "delay"
)

type Airline struct {
	Id               int         `json:"id"`
	From             string      `json:"from"`
	To               string      `json:"to"`
	Date             time.Time   `json:"date"`
	AirportName      string      `json:"airport_name"`
	NoOfSeatsInClass Class       `json:"noOfSeatInclass"`
	Name             string      `json:"name"`
	FlightNumber     string      `json:"flight_number"`
	DepartureTime    string      `json:"departure_time"`
	ArrivalTime      string      `json:"arrival_time"`
	Terminal         string      `json:"terminal"`
	Gate             string      `json:"gate"`
	BasePrices       ClassPrices `json:"base_prices"`
	FlightDuration   string      `json:"flight_duration"`
	FlightStatus     string      `json:"flight_status"` // scheduled, boarding, departed, delayed, etc.
	IsCodeShare      bool        `json:"is_code_share"`
}
type Airport struct {
	Code      string   `json:"code"` //IATA code
	Name      string   `json:"name"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Timezone  string   `json:"timezone"`
	Terminals []string `json:"terminals"`
}
