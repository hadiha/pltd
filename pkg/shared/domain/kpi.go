package domain

import (
	"time"
)

// Kpi model
type Kpi struct {
	ID         string    `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	Field      string    `gorm:"column:field;type:varchar(255)" json:"field"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of Kpi model
func (Kpi) TableName() string {
	return "kpis"
}

