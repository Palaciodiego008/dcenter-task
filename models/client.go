package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"column:name"`
	LastName      string `gorm:"column:last_name"`
	StreetAddress string `gorm:"column:street_address"`
	Phone         string `gorm:"column:phone"`
	Email         string `gorm:"column:email"`
}

// TableName specifies the table name for the Client model
func (Client) TableName() string {
	return "clients"
}
