package entities

import (
	"time"

	"gorm.io/gorm"
)

type Quotation struct {
	gorm.Model
	SoldToCode    string
	ShipToCode    string
	DocumentCode  string
	QuotationCode string
	ProjectName   string
	DeliveryDate  time.Time
	PlantCode     string
	DiscountType  string
	Status        string
	Items         []QuotationItem `gorm:"foreignKey:QuotationId"`
}

type QuotationItem struct {
	gorm.Model
	ItemIndex   int16
	QuotationId uint
	ItemType    string
	Price       float64
	Amount      float32
}
