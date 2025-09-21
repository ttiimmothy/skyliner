package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"skyliner/internal/config"
	"skyliner/internal/db/models"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	billingportal "github.com/stripe/stripe-go/v76/billingportal/session"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/webhook"
	"gorm.io/gorm"
)

type PaymentHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewPaymentHandler(db *gorm.DB, cfg *config.Config) *PaymentHandler {
	// Set Stripe secret key
	stripe.Key = cfg.StripeSecretKey
	return &PaymentHandler{db: db, cfg: cfg}
}

type CheckoutSessionRequest struct {
	BookingID uint `json:"booking_id" binding:"required"`
}

type CheckoutSessionResponse struct {
	SessionID string `json:"session_id"`
	URL       string `json:"url"`
}

type BillingPortalRequest struct {
	BookingID uint `json:"booking_id" binding:"required"`
}

type BillingPortalResponse struct {
	URL string `json:"url"`
}

func (h *PaymentHandler) CreateCheckoutSession(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req CheckoutSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get booking
	var booking models.Booking
	if err := h.db.Preload("User").Where("id = ? AND user_id = ?", req.BookingID, userID).First(&booking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}

	if booking.Status != models.StatusHold {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking is not in hold status"})
		return
	}

	// Create or get Stripe customer
	customerID, err := h.getOrCreateStripeCustomer(&booking.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	// Create checkout session
	params := &stripe.CheckoutSessionParams{
		Customer:           stripe.String(customerID),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(booking.Currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(fmt.Sprintf("Flight Booking - %s", booking.PNR)),
					},
					UnitAmount: stripe.Int64(int64(booking.TotalAmount * 100)), // Convert to cents
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("http://localhost:5193/booking/" + strconv.Itoa(int(booking.ID)) + "?success=true"),
		CancelURL:  stripe.String("http://localhost:5193/booking/" + strconv.Itoa(int(booking.ID)) + "?cancelled=true"),
		Metadata: map[string]string{
			"booking_id": strconv.Itoa(int(booking.ID)),
		},
	}

	session, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout session"})
		return
	}

	// Update booking with session ID
	if err := h.db.Model(&booking).Update("stripe_session_id", session.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}

	c.JSON(http.StatusOK, CheckoutSessionResponse{
		SessionID: session.ID,
		URL:       session.URL,
	})
}

func (h *PaymentHandler) CreateBillingPortal(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req BillingPortalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get booking
	var booking models.Booking
	if err := h.db.Preload("User").Where("id = ? AND user_id = ?", req.BookingID, userID).First(&booking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}

	// Get or create Stripe customer
	customerID, err := h.getOrCreateStripeCustomer(&booking.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	// Create billing portal session
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(customerID),
		ReturnURL: stripe.String("http://localhost:5193/billing"),
	}

	session, err := billingportal.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create billing portal session"})
		return
	}

	c.JSON(http.StatusOK, BillingPortalResponse{
		URL: session.URL,
	})
}

func (h *PaymentHandler) HandleStripeWebhook(c *gin.Context) {
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	// Verify webhook signature
	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), h.cfg.StripeWebhookSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signature"})
		return
	}

	// Handle the event
	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session data"})
			return
		}
		h.handleCheckoutSessionCompleted(&session)
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment intent data"})
			return
		}
		h.handlePaymentIntentSucceeded(&paymentIntent)
	default:
		fmt.Printf("Unhandled event type: %s\n", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *PaymentHandler) getOrCreateStripeCustomer(user *models.User) (string, error) {
	// Check if user already has a Stripe customer ID stored
	// For now, create a new customer each time
	params := &stripe.CustomerParams{
		Email: stripe.String(user.Email),
		Name:  stripe.String(user.FirstName + " " + user.LastName),
	}

	customer, err := customer.New(params)
	if err != nil {
		return "", err
	}

	return customer.ID, nil
}

func (h *PaymentHandler) handleCheckoutSessionCompleted(session *stripe.CheckoutSession) {
	// Get booking ID from metadata
	bookingIDStr, exists := session.Metadata["booking_id"]
	if !exists {
		fmt.Printf("No booking_id in session metadata\n")
		return
	}

	bookingID, err := strconv.ParseUint(bookingIDStr, 10, 32)
	if err != nil {
		fmt.Printf("Invalid booking_id in session metadata: %s\n", bookingIDStr)
		return
	}

	// Update booking status to paid
	if err := h.db.Model(&models.Booking{}).Where("id = ?", uint(bookingID)).Update("status", models.StatusPaid).Error; err != nil {
		fmt.Printf("Failed to update booking status: %v\n", err)
		return
	}

	// Create payment record
	payment := models.Payment{
		BookingID:       uint(bookingID),
		StripePaymentID: session.PaymentIntent.ID,
		Amount:          float64(session.AmountTotal) / 100, // Convert from cents
		Currency:        string(session.Currency),
		Status:          "succeeded",
	}

	if err := h.db.Create(&payment).Error; err != nil {
		fmt.Printf("Failed to create payment record: %v\n", err)
		return
	}

	fmt.Printf("Booking %d marked as paid\n", bookingID)
}

func (h *PaymentHandler) handlePaymentIntentSucceeded(paymentIntent *stripe.PaymentIntent) {
	// This is handled by checkout.session.completed for now
	fmt.Printf("Payment intent succeeded: %s\n", paymentIntent.ID)
}
