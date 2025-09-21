package ws

import (
	"encoding/json"
	"strconv"
	"time"
)

type EventType string

const (
	EventPriceTick     EventType = "priceTick"
	EventSeatUpdate    EventType = "seatUpdate"
	EventBookingStatus EventType = "bookingStatus"
)

type Event struct {
	Type      EventType   `json:"type"`
	Channel   string      `json:"channel"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

type PriceTickData struct {
	SearchHash string  `json:"search_hash"`
	FlightID   uint    `json:"flight_id"`
	FareID     uint    `json:"fare_id"`
	OldPrice   float64 `json:"old_price"`
	NewPrice   float64 `json:"new_price"`
	Currency   string  `json:"currency"`
}

type SeatUpdateData struct {
	FlightID uint   `json:"flight_id"`
	SeatID   uint   `json:"seat_id"`
	Status   string `json:"status"`
	Row      int    `json:"row"`
	Column   string `json:"column"`
	Class    string `json:"class"`
}

type BookingStatusData struct {
	BookingID uint   `json:"booking_id"`
	PNR       string `json:"pnr"`
	Status    string `json:"status"`
	Message   string `json:"message,omitempty"`
}

func NewPriceTickEvent(searchHash string, flightID, fareID uint, oldPrice, newPrice float64, currency string) *Event {
	return &Event{
		Type:    EventPriceTick,
		Channel: "priceTick:" + searchHash,
		Data: PriceTickData{
			SearchHash: searchHash,
			FlightID:   flightID,
			FareID:     fareID,
			OldPrice:   oldPrice,
			NewPrice:   newPrice,
			Currency:   currency,
		},
		Timestamp: time.Now(),
	}
}

func NewSeatUpdateEvent(flightID, seatID uint, status, row, column, class string) *Event {
	return &Event{
		Type:    EventSeatUpdate,
		Channel: "seatUpdate:" + string(rune(flightID)),
		Data: SeatUpdateData{
			FlightID: flightID,
			SeatID:   seatID,
			Status:   status,
			Row:      func() int { r, _ := strconv.Atoi(row); return r }(),
			Column:   column,
			Class:    class,
		},
		Timestamp: time.Now(),
	}
}

func NewBookingStatusEvent(bookingID uint, pnr, status, message string) *Event {
	return &Event{
		Type:    EventBookingStatus,
		Channel: "bookingStatus:" + string(rune(bookingID)),
		Data: BookingStatusData{
			BookingID: bookingID,
			PNR:       pnr,
			Status:    status,
			Message:   message,
		},
		Timestamp: time.Now(),
	}
}

func (e *Event) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}
