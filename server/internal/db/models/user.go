package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleTraveler Role = "traveler"
	RoleAgent    Role = "agent"
	RoleAdmin    Role = "admin"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Email        string         `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Role         Role           `json:"role" gorm:"default:traveler"`
	GoogleID     *string        `json:"-" gorm:"uniqueIndex"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	Bookings []Booking `json:"bookings,omitempty" gorm:"foreignKey:UserID"`
}
