package domain

import (
	"time"
)

// Business model
type BusinessReport struct {
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

// type BusinessSum struct {
// 	ID        string `json:"id"`
// 	Day 			string `json:"day"`
// 	Date 			string `json:"date"`
// 	Dmn 			float32 `json:"dmn"`
// 	Cad 			float32 `json:"cad"`
// 	Dmp 			float32 `json:"dmp"`
// 	Bp 				float32 `json:"bp"`
// 	Cad_mw 		float32 `json:"cad_mw"`
// 	Po				float32 `json:"po"`
// 	Mo				float32 `json:"mo"`
// 	Fd				float32 `json:"fd"`
// 	Total			float32 `json:"total"`
// }

// TableName return table name of Business model
func (BusinessReport) TableName() string {
	return "pltd_daily"
}



type Business struct {
	ID         string    `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	Tgl_input  string    `gorm:"column:tgl_input;type:varchar(255)" json:"tgl_input"`
	Tgl_laporan  string    `gorm:"column:tgl_laporan;type:varchar(255)" json:"tgl_laporan"`
	Jenis_pembangkit  string    `gorm:"column:jenis_pembangkit;type:varchar(255)" json:"jenis_pembangkit"`
	Shift  string    `gorm:"column:shift;type:varchar(255)" json:"shift"`
	Sh_ipp_ap  string    `gorm:"column:sh_ipp_ap;type:varchar(255)" json:"sh_ipp_ap"`
	Sistem  string    `gorm:"column:sistem;type:varchar(255)" json:"sistem"`
	Isolated  string    `gorm:"column:isolated;type:varchar(255)" json:"isolated"`
	Sentral  string    `gorm:"column:sentral;type:varchar(255)" json:"sentral"`
	Unit_mesin  string    `gorm:"column:unit_mesin;type:varchar(255)" json:"unit_mesin"`
	Tegangan  string    `gorm:"column:tegangan;type:varchar(255)" json:"tegangan"`
	Daya_pasang  string    `gorm:"column:daya_pasang;type:varchar(255)" json:"daya_pasang"`
	Daya_mampu  string    `gorm:"column:daya_mampu;type:varchar(255)" json:"daya_mampu"`
	Dmp  string    `gorm:"column:dmp;type:varchar(255)" json:"dmp"`
	Beban  string    `gorm:"column:beban;type:varchar(255)" json:"beban"`
	Status  string    `gorm:"column:status;type:varchar(255)" json:"status"`
	Deskripsi  string    `gorm:"column:deskripsi;type:varchar(255)" json:"deskripsi"`
	Tgl_mulai  string    `gorm:"column:tgl_mulai;type:varchar(255)" json:"tgl_mulai"`
	Tgl_selesai  string    `gorm:"column:tgl_selesai;type:varchar(255)" json:"tgl_selesai"`
	Progress  string    `gorm:"column:progress;type:varchar(255)" json:"progress"`
	Outage  string    `gorm:"column:outage;type:varchar(255)" json:"outage"`
	Derating  string    `gorm:"column:derating;type:varchar(255)" json:"derating"`
	Hop_bb  string    `gorm:"column:hop_bb;type:varchar(255)" json:"hop_bb"`
	Vol_bb  string    `gorm:"column:vol_bb;type:varchar(255)" json:"vol_bb"`
	Vol_bbm  string    `gorm:"column:vol_bbm;type:varchar(255)" json:"vol_bbm"`
	CreatedAt  time.Time    `gorm:"column:created_at;type:varchar(255)" json:"created_at"`
	UpdatedAt  time.Time    `gorm:"column:updated_at;type:varchar(255)" json:"updated_at"`
	Jumlah  string    `gorm:"column:jumlah;type:varchar(255)" json:"jumlah"`
}

// TableName return table name of Business model
func (Business) TableName() string {
	return "pltd_pembangkit"
}
