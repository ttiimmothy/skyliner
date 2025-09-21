package models

import (
	"time"
)

type Airport struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code" gorm:"uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"not null"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	DepartureFlights []Flight `json:"departure_flights,omitempty" gorm:"foreignKey:OriginID"`
	ArrivalFlights   []Flight `json:"arrival_flights,omitempty" gorm:"foreignKey:DestinationID"`
}

type Airline struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code" gorm:"uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Flights []Flight `json:"flights,omitempty" gorm:"foreignKey:AirlineID"`
}

type Flight struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Number        string    `json:"number" gorm:"not null"`
	AirlineID     uint      `json:"airline_id" gorm:"not null"`
	OriginID      uint      `json:"origin_id" gorm:"not null"`
	DestinationID uint      `json:"destination_id" gorm:"not null"`
	DepartureTime time.Time `json:"departure_time" gorm:"not null"`
	ArrivalTime   time.Time `json:"arrival_time" gorm:"not null"`
	Duration      int       `json:"duration" gorm:"not null"` // in minutes
	Stops         int       `json:"stops" gorm:"default:0"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// Relations
	Airline     Airline   `json:"airline" gorm:"foreignKey:AirlineID"`
	Origin      Airport   `json:"origin" gorm:"foreignKey:OriginID"`
	Destination Airport   `json:"destination" gorm:"foreignKey:DestinationID"`
	Fares       []Fare    `json:"fares,omitempty" gorm:"foreignKey:FlightID"`
	SeatMaps    []SeatMap `json:"seat_maps,omitempty" gorm:"foreignKey:FlightID"`
	Segments    []Segment `json:"segments,omitempty" gorm:"foreignKey:FlightID"`
}

type Fare struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FlightID  uint      `json:"flight_id" gorm:"not null"`
	Class     string    `json:"class" gorm:"not null"`     // economy, business, first
	FareType  string    `json:"fare_type" gorm:"not null"` // basic, standard, flexible
	BasePrice float64   `json:"base_price" gorm:"not null"`
	Currency  string    `json:"currency" gorm:"default:USD"`
	Available int       `json:"available" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Flight Flight `json:"flight" gorm:"foreignKey:FlightID"`
}
