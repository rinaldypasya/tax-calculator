package models

import "time"

type Taxes struct {
	ID             uint   `gorm:"primary_key"`
	Code           int64 `gorm:"column:code" json:"code"`
	Name           string `gorm:"column:name" json:"name"`
	IsRefundable   bool   `gorm:"column:is_refundable" json:"is_refundable"`
	Configurations []TaxConfigurations
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}

func (Taxes) TableName() string {
	return "taxes"
}
