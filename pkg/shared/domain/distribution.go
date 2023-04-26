package domain

import (
	"time"
)

// Distribution model
type Distribution struct {
	ID    string    `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	Instansi 	string	 `gorm:"column:instansi;type:varchar(15)" json:"instansi"` 
	Pembangkit 	string	 `gorm:"column:pembangkit;type:varchar(50)" json:"pembangkit"` 
	Provinsi 			string `gorm:"column:provinsi;type:varchar(50)" json:"provinsi"`
	Jenis_kit				string `gorm:"column:jenis_kit;type:varchar(50)" json:"jenis_kit"`
	Bbm 		string `gorm:"column:bbm;type:varchar(50)" json:"bbm"`
	Daya_pasang 		int32 `gorm:"column:daya_pasang;type:integer" json:"daya_pasang"`
	Daya_mampu 		int32 `gorm:"column:daya_mampu;type:integer" json:"daya_mampu"`
	Tahun 		string `gorm:"column:tahun;type:varchar(50)" json:"tahun"`
	Kondisi 		string `gorm:"column:kondisi;type:varchar(50)" json:"kondisi"`
	Lattitude 		string `gorm:"column:lattitude;type:varchar(50)" json:"lattitude"`
	Longitude 		string `gorm:"column:longitude;type:varchar(50)" json:"longitude"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type DistributionSum struct {
	Daya_pasang 		int32 `gorm:"column:daya_pasang;type:integer" json:"daya_pasang"`
	Daya_mampu 		int32 `gorm:"column:daya_mampu;type:integer" json:"daya_mampu"`
	Sentral 	int32	 `json:"sentral"` 
	Mesin 	int32	 `json:"mesin"` 
}

// TableName return table name of Distribution model
func (Distribution) TableName() string {
	return "pltd_sebaran"
}

func (DistributionSum) TableName() string {
	return "pltd_sebaran"
}

