package handlers

import (
	"net/http"
	"strconv"
	"time"

	"skyliner/internal/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SearchHandler struct {
	db *gorm.DB
}

func NewSearchHandler(db *gorm.DB) *SearchHandler {
	return &SearchHandler{db: db}
}

type SearchRequest struct {
	TripType    string `json:"trip_type" binding:"required"` // one-way, round-trip, multi-city
	Legs        []Leg  `json:"legs" binding:"required"`
	Passengers  int    `json:"passengers" binding:"required,min=1,max=9"`
	Cabin       string `json:"cabin"`       // economy, business, first
	Flexibility int    `json:"flexibility"` // days
}

type Leg struct {
	Origin      string `json:"origin" binding:"required"`
	Destination string `json:"destination" binding:"required"`
	Date        string `json:"date" binding:"required"`
}

type SearchResponse struct {
	Flights []FlightResult `json:"flights"`
	Total   int            `json:"total"`
}

type FlightResult struct {
	ID            uint           `json:"id"`
	Number        string         `json:"number"`
	Airline       models.Airline `json:"airline"`
	Origin        models.Airport `json:"origin"`
	Destination   models.Airport `json:"destination"`
	DepartureTime time.Time      `json:"departure_time"`
	ArrivalTime   time.Time      `json:"arrival_time"`
	Duration      int            `json:"duration"`
	Stops         int            `json:"stops"`
	Fares         []FareResult   `json:"fares"`
}

type FareResult struct {
	ID        uint    `json:"id"`
	Class     string  `json:"class"`
	FareType  string  `json:"fare_type"`
	BasePrice float64 `json:"base_price"`
	Currency  string  `json:"currency"`
	Available int     `json:"available"`
}

func (h *SearchHandler) GetAirports(c *gin.Context) {
	var airports []models.Airport
	if err := h.db.Find(&airports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch airports"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"airports": airports})
}

func (h *SearchHandler) GetAirlines(c *gin.Context) {
	var airlines []models.Airline
	if err := h.db.Find(&airlines).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch airlines"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"airlines": airlines})
}

func (h *SearchHandler) SearchFlights(c *gin.Context) {
	var req SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// For now, search the first leg only
	if len(req.Legs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one leg required"})
		return
	}

	leg := req.Legs[0]
	departureDate, err := time.Parse("2006-01-02", leg.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	// Build query
	query := h.db.Model(&models.Flight{}).
		Preload("Airline").
		Preload("Origin").
		Preload("Destination").
		Preload("Fares").
		Where("origin_id IN (SELECT id FROM airports WHERE code = ?)", leg.Origin).
		Where("destination_id IN (SELECT id FROM airports WHERE code = ?)", leg.Destination)

	// Add date range for flexibility
	if req.Flexibility > 0 {
		startDate := departureDate.AddDate(0, 0, -req.Flexibility)
		endDate := departureDate.AddDate(0, 0, req.Flexibility)
		query = query.Where("departure_time BETWEEN ? AND ?", startDate, endDate)
	} else {
		startDate := departureDate.Truncate(24 * time.Hour)
		endDate := startDate.Add(24 * time.Hour)
		query = query.Where("departure_time BETWEEN ? AND ?", startDate, endDate)
	}

	// Filter by cabin class if specified
	if req.Cabin != "" {
		query = query.Joins("JOIN fares ON fares.flight_id = flights.id").
			Where("fares.class = ?", req.Cabin)
	}

	var flights []models.Flight
	if err := query.Find(&flights).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search flights"})
		return
	}

	// Convert to response format
	var results []FlightResult
	for _, flight := range flights {
		var fares []FareResult
		for _, fare := range flight.Fares {
			fares = append(fares, FareResult{
				ID:        fare.ID,
				Class:     fare.Class,
				FareType:  fare.FareType,
				BasePrice: fare.BasePrice,
				Currency:  fare.Currency,
				Available: fare.Available,
			})
		}

		results = append(results, FlightResult{
			ID:            flight.ID,
			Number:        flight.Number,
			Airline:       flight.Airline,
			Origin:        flight.Origin,
			Destination:   flight.Destination,
			DepartureTime: flight.DepartureTime,
			ArrivalTime:   flight.ArrivalTime,
			Duration:      flight.Duration,
			Stops:         flight.Stops,
			Fares:         fares,
		})
	}

	c.JSON(http.StatusOK, SearchResponse{
		Flights: results,
		Total:   len(results),
	})
}

func (h *SearchHandler) GetSeatMap(c *gin.Context) {
	flightIDStr := c.Param("id")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flight ID"})
		return
	}

	var seatMaps []models.SeatMap
	if err := h.db.Preload("Seats").Where("flight_id = ?", uint(flightID)).Find(&seatMaps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch seat map"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"seat_maps": seatMaps})
}
