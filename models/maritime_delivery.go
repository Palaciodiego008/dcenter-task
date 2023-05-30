package models

import (
	"time"

	"gorm.io/gorm"
)

type TruckDelivery struct {
	gorm.Model
	ID               uint      `gorm:"primaryKey"`
	ClientID         string    `json:"client_id" validate:"required"`
	ProductType      string    `json:"product_type" validate:"required"`
	Quantity         int       `json:"quantity" validate:"required,gte=1"`
	RegistrationDate time.Time `json:"registration_date" validate:"required,datetime"`
	DeliveryDate     time.Time `json:"delivery_date" validate:"required,datetime"`
	Warehouse        string    `json:"warehouse" validate:"required"`
	ShippingPrice    float64   `json:"shipping_price" validate:"required,gte=0"`
	DiscountedPrice  float64   `json:"discounted_price" validate:"required,gte=0"`

	VehiclePlate string `json:"vehicle_plate" validate:"required,regexp=^[A-Z]{3}[0-9]{3}$"`
	GuideNumber  string `json:"guide_number" validate:"required,len=10"`
}
