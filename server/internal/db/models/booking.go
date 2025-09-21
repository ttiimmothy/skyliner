package models

import (
	"time"
)

type BookingStatus string

const (
	StatusHold      BookingStatus = "hold"
	StatusPaid      BookingStatus = "paid"
	StatusTicketed  BookingStatus = "ticketed"
	StatusCancelled BookingStatus = "cancelled"
)

type Booking struct {
	ID              uint          `json:"id" gorm:"primaryKey"`
	PNR             string        `json:"pnr" gorm:"uniqueIndex;not null"`
	UserID          uint          `json:"user_id" gorm:"not null"`
	Status          BookingStatus `json:"status" gorm:"default:hold"`
	TotalAmount     float64       `json:"total_amount" gorm:"not null"`
	Currency        string        `json:"currency" gorm:"default:USD"`
	StripeSessionID *string       `json:"stripe_session_id"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`

	// Relations
	User       User        `json:"user" gorm:"-"`
	Itinerary  Itinerary   `json:"itinerary" gorm:"-"`
	Passengers []Passenger `json:"passengers,omitempty" gorm:"-"`
	Payments   []Payment   `json:"payments,omitempty" gorm:"-"`
}

type Itinerary struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BookingID uint      `json:"booking_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Segments []Segment `json:"segments,omitempty" gorm:"-"`
}

type Segment struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ItineraryID uint      `json:"itinerary_id" gorm:"not null"`
	FlightID    uint      `json:"flight_id" gorm:"not null"`
	FareID      uint      `json:"fare_id" gorm:"not null"`
	SeatNumber  *string   `json:"seat_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relations
	Itinerary Itinerary `json:"itinerary" gorm:"-"`
	Flight    Flight    `json:"flight" gorm:"-"`
	Fare      Fare      `json:"fare" gorm:"-"`
}

type Passenger struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	BookingID   uint       `json:"booking_id" gorm:"not null"`
	FirstName   string     `json:"first_name" gorm:"not null"`
	LastName    string     `json:"last_name" gorm:"not null"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Passport    *string    `json:"passport"`
	SSR         *string    `json:"ssr"` // Special Service Request
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	// Relations
	Booking Booking `json:"booking" gorm:"-"`
}

type Payment struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	BookingID       uint      `json:"booking_id" gorm:"not null"`
	StripePaymentID string    `json:"stripe_payment_id" gorm:"not null"`
	Amount          float64   `json:"amount" gorm:"not null"`
	Currency        string    `json:"currency" gorm:"default:USD"`
	Status          string    `json:"status" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	// Relations
	Booking Booking `json:"booking" gorm:"-"`
}
