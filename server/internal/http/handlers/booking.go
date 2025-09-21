package handlers

import (
	"net/http"
	"strconv"
	"time"

	"skyliner/internal/config"
	"skyliner/internal/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookingHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewBookingHandler(db *gorm.DB, cfg *config.Config) *BookingHandler {
	return &BookingHandler{db: db, cfg: cfg}
}

type CreateBookingRequest struct {
	Segments   []SegmentRequest   `json:"segments" binding:"required"`
	Passengers []PassengerRequest `json:"passengers" binding:"required"`
	Seats      []SeatRequest      `json:"seats"`
	Extras     []ExtraRequest     `json:"extras"`
}

type SegmentRequest struct {
	FlightID   uint    `json:"flight_id" binding:"required"`
	FareID     uint    `json:"fare_id" binding:"required"`
	SeatNumber *string `json:"seat_number"`
}

type PassengerRequest struct {
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Passport    *string    `json:"passport"`
	SSR         *string    `json:"ssr"`
}

type SeatRequest struct {
	SegmentID uint `json:"segment_id" binding:"required"`
	SeatID    uint `json:"seat_id" binding:"required"`
}

type ExtraRequest struct {
	Type  string  `json:"type" binding:"required"` // baggage, meal, etc.
	Price float64 `json:"price" binding:"required"`
}

type BookingResponse struct {
	Booking     models.Booking `json:"booking"`
	TotalAmount float64        `json:"total_amount"`
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Start transaction
	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Generate PNR
	pnr := h.generatePNR()

	// Calculate total amount
	totalAmount := 0.0
	for _, segment := range req.Segments {
		var fare models.Fare
		if err := tx.First(&fare, segment.FareID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fare ID"})
			return
		}
		totalAmount += fare.BasePrice
	}

	// Add extras
	for _, extra := range req.Extras {
		totalAmount += extra.Price
	}

	// Create booking
	booking := models.Booking{
		PNR:         pnr,
		UserID:      userID,
		Status:      models.StatusHold,
		TotalAmount: totalAmount,
		Currency:    "USD",
	}

	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	// Create itinerary
	itinerary := models.Itinerary{
		BookingID: booking.ID,
	}
	if err := tx.Create(&itinerary).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create itinerary"})
		return
	}

	// Create segments
	for _, segmentReq := range req.Segments {
		segment := models.Segment{
			ItineraryID: itinerary.ID,
			FlightID:    segmentReq.FlightID,
			FareID:      segmentReq.FareID,
			SeatNumber:  segmentReq.SeatNumber,
		}
		if err := tx.Create(&segment).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create segment"})
			return
		}
	}

	// Create passengers
	for _, passengerReq := range req.Passengers {
		passenger := models.Passenger{
			BookingID:   booking.ID,
			FirstName:   passengerReq.FirstName,
			LastName:    passengerReq.LastName,
			Email:       passengerReq.Email,
			Phone:       passengerReq.Phone,
			DateOfBirth: passengerReq.DateOfBirth,
			Passport:    passengerReq.Passport,
			SSR:         passengerReq.SSR,
		}
		if err := tx.Create(&passenger).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create passenger"})
			return
		}
	}

	// Create baggage for extras
	for _, extra := range req.Extras {
		if extra.Type == "baggage" {
			baggage := models.Baggage{
				BookingID: booking.ID,
				Type:      "checked",
				Price:     extra.Price,
			}
			if err := tx.Create(&baggage).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create baggage"})
				return
			}
		}
	}

	// Update seat availability
	for _, seatReq := range req.Seats {
		if err := tx.Model(&models.Seat{}).Where("id = ?", seatReq.SeatID).Update("status", models.SeatSelected).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update seat status"})
			return
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit booking"})
		return
	}

	// Load booking with relations
	if err := h.db.Preload("User").Preload("Itinerary.Segments.Flight").Preload("Itinerary.Segments.Fare").Preload("Passengers").First(&booking, booking.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load booking"})
		return
	}

	c.JSON(http.StatusCreated, BookingResponse{
		Booking:     booking,
		TotalAmount: totalAmount,
	})
}

func (h *BookingHandler) GetBooking(c *gin.Context) {
	bookingIDStr := c.Param("id")
	bookingID, err := strconv.ParseUint(bookingIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	userID := c.GetUint("user_id")

	var booking models.Booking
	if err := h.db.Preload("User").Preload("Itinerary.Segments.Flight.Airline").Preload("Itinerary.Segments.Flight.Origin").Preload("Itinerary.Segments.Flight.Destination").Preload("Itinerary.Segments.Fare").Preload("Passengers").Preload("Payments").Where("id = ? AND user_id = ?", uint(bookingID), userID).First(&booking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"booking": booking})
}

func (h *BookingHandler) IssueBooking(c *gin.Context) {
	bookingIDStr := c.Param("id")
	bookingID, err := strconv.ParseUint(bookingIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	userID := c.GetUint("user_id")

	var booking models.Booking
	if err := h.db.Where("id = ? AND user_id = ?", uint(bookingID), userID).First(&booking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}

	if booking.Status != models.StatusPaid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking must be paid before issuing"})
		return
	}

	// Update booking status to ticketed
	if err := h.db.Model(&booking).Update("status", models.StatusTicketed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to issue booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking issued successfully"})
}

func (h *BookingHandler) CancelBooking(c *gin.Context) {
	bookingIDStr := c.Param("id")
	bookingID, err := strconv.ParseUint(bookingIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	userID := c.GetUint("user_id")

	var booking models.Booking
	if err := h.db.Where("id = ? AND user_id = ?", uint(bookingID), userID).First(&booking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}

	if booking.Status == models.StatusCancelled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking already cancelled"})
		return
	}

	// Update booking status to cancelled
	if err := h.db.Model(&booking).Update("status", models.StatusCancelled).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}

func (h *BookingHandler) GetAllBookings(c *gin.Context) {
	// Admin/Agent only
	var bookings []models.Booking
	if err := h.db.Preload("User").Preload("Itinerary.Segments.Flight").Preload("Passengers").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func (h *BookingHandler) WaiveBooking(c *gin.Context) {
	// TODO: Implement waiver logic
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Waiver not implemented yet"})
}

func (h *BookingHandler) RepriceBookings(c *gin.Context) {
	// TODO: Implement reprice logic
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Reprice not implemented yet"})
}

func (h *BookingHandler) generatePNR() string {
	// Simple PNR generation - in production, use a more sophisticated method
	return "SKY" + strconv.FormatInt(time.Now().Unix(), 36)
}
