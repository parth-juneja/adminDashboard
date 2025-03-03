package models

import (
	"time"

	"gorm.io/gorm"
)

// Enum for SubscriptionType
type SubscriptionType string

const (
	OneMonth     SubscriptionType = "1_month"
	ThreeMonths  SubscriptionType = "3_months"
	SixMonths    SubscriptionType = "6_months"
	TwelveMonths SubscriptionType = "12_months"
)

// Customer model
type Customer struct {
	gorm.Model
	Name          string           `gorm:"not null" json:"name"`
	StartDate     time.Time        `json:"start_date"`
	RemainingDays int              `json:"remaining_days"`
	Subscription  SubscriptionType `gorm:"type:varchar(10)" json:"subscription"`
	PaymentAmount float64          `json:"payment_amount"`
	LastPaid      time.Time        `json:"last_paid"`
}
