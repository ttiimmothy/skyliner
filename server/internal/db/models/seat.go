package models

import (
	"time"
)

type SeatMap struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FlightID  uint      `json:"flight_id" gorm:"not null"`
	Class     string    `json:"class" gorm:"not null"` // economy, business, first
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Flight Flight `json:"flight" gorm:"foreignKey:FlightID"`
	Seats  []Seat `json:"seats,omitempty" gorm:"foreignKey:SeatMapID"`
}

type SeatStatus string

const (
	SeatAvailable SeatStatus = "available"
	SeatOccupied  SeatStatus = "occupied"
	SeatBlocked   SeatStatus = "blocked"
	SeatSelected  SeatStatus = "selected"
)

type Seat struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	SeatMapID uint       `json:"seat_map_id" gorm:"not null"`
	Row       int        `json:"row" gorm:"not null"`
	Column    string     `json:"column" gorm:"not null"`
	Class     string     `json:"class" gorm:"not null"`
	Status    SeatStatus `json:"status" gorm:"default:available"`
	Price     *float64   `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Relations
	SeatMap SeatMap `json:"seat_map" gorm:"foreignKey:SeatMapID"`
}

type Baggage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BookingID uint      `json:"booking_id" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"` // checked, carry-on
	Weight    int       `json:"weight"`               // in kg
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Booking Booking `json:"booking" gorm:"foreignKey:BookingID"`
}
