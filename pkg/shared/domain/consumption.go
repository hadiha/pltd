package domain

import (
	"time"
)

// Consumption model
type Consumption struct {
	ID    string    `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	Unit 	string	 `gorm:"column:unit;type:varchar(15)" json:"unit"` 
	Pembangkit 	string	 `gorm:"column:pembangkit;type:varchar(50)" json:"pembangkit"` 
	Minyak 	int32	 `gorm:"column:minyak;type:integer" json:"minyak"` 
	Gas 	int32	 `gorm:"column:gas;type:integer" json:"gas"` 
	Batubara 	int32	 `gorm:"column:batubara;type:integer" json:"batubara"` 
	Biomass 	int32	 `gorm:"column:biomass;type:integer" json:"biomass"` 
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of Consumption model
func (Consumption) TableName() string {
	return "pltd_bahan_bakar"
}

