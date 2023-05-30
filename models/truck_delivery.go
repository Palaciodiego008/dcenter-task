package models

import (
	"time"

	"gorm.io/gorm"
)

// Struct for maritime deliveries
type ShipDelivery struct {
	gorm.Model
	ID               uint      `gorm:"primaryKey"`
	ClientID         int       `json:"client_id" validate:"required"`
	ProductType      string    `json:"product_type" validate:"required"`
	Quantity         int       `json:"quantity" validate:"required,gte=1"`
	RegistrationDate time.Time `json:"registration_date" validate:"required,datetime"`
	DeliveryDate     time.Time `json:"delivery_date" validate:"required,datetime"`
	Port             string    `json:"port" validate:"required"`
	ShippingPrice    float64   `json:"shipping_price" validate:"required,gte=0"`
	DiscountedPrice  float64   `json:"discounted_price" validate:"required,gte=0"`
	FleetNumber      string    `json:"fleet_number" validate:"required,len=10"`
	GuideNumber      string    `json:"guide_number" validate:"required,len=10"`
}

// TableName specifies the table name for the ShipDelivery model
func (ShipDelivery) TableName() string {
	return "ship_deliveries"
}
