package domain

import (
	"time"
)

// Monitoring model
type Monitoring struct {
	ID         string    `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	Unit 	string	 `gorm:"column:unit;type:varchar(255)" json:"unit"` 
	Day 	string	 `gorm:"column:day;type:varchar(255)" json:"day"`	 
	Date 	string	 `gorm:"column:date;type:varchar(255)" json:"date"` 
	Dmn 	float32	 `gorm:"column:dmn;type:double" json:"dmn"` 
	Cad 	float32	 `gorm:"column:cad;type:double" json:"cad"` 
	Dmp 	float32	 `gorm:"column:dmp;type:double" json:"dmp"` 
	Bp 	float32	 `gorm:"column:bp;type:double" json:"bp"` 
	Cad_mw 	float32	 `gorm:"column:cad_mw;type:double" json:"cad_mw"` 
	Status	string			 `gorm:"column:status;type:varchar(255)" json:"status"` 
	Description	string		 `gorm:"column:description;type:varchar(255)" json:"description"` 
	Po	float32	 `gorm:"column:po;type:double" json:"po"` 
	Mo	float32	 `gorm:"column:mo;type:double" json:"mo"` 
	Fd	float32	 `gorm:"column:fd;type:double" json:"fd"` 
	Total	float32	 `gorm:"column:total;type:double" json:"total"` 
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of Monitoring model
func (Monitoring) TableName() string {
	return "pltd_daily"
}

