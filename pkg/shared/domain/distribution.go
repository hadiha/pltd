package domain

import "time"

// Distribution model
type Distribution struct {
	ID         string    `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	NamaRegional  string    `gorm:"column:nama_regional;type:varchar(255)" json:"nama_regional"`
	NomorRegional  string    `gorm:"column:nomor_regional;type:varchar(255)" json:"nomor_regional"`
	NamaSistem  string    `gorm:"column:nama_sistem;type:varchar(255)" json:"nama_sistem"`
	KodeSistem  string    `gorm:"column:kode_sistem;type:varchar(255)" json:"kode_sistem"`
	NamaUnit  string    `gorm:"column:nama_unit;type:varchar(255)" json:"nama_unit"`
	KodeUnit  string    `gorm:"column:kode_unit;type:varchar(255)" json:"kode_unit"`
	NamaPembangkit  string    `gorm:"column:nama_pembangkit;type:varchar(255)" json:"nama_pembangkit"`
	NomorPembangkit  string    `gorm:"column:nomor_pembangkit;type:varchar(255)" json:"nomor_pembangkit"`
	Instansi  string    `gorm:"column:instansi;type:varchar(255)" json:"instansi"`
	Provinsi  string    `gorm:"column:provinsi;type:varchar(255)" json:"provinsi"`
	Alamat  string    `gorm:"column:alamat;type:varchar(255)" json:"alamat"`
	JenisKit  string    `gorm:"column:jenis_kit;type:varchar(255)" json:"jenis_kit"`
	Bbm  string    `gorm:"column:bbm;type:varchar(255)" json:"bbm"`
	DayaPasang  string    `gorm:"column:daya_pasang;type:varchar(255)" json:"daya_pasang"`
	DayaMampu  string    `gorm:"column:daya_mampu;type:varchar(255)" json:"daya_mampu"`
	Tahun  string    `gorm:"column:tahun;type:varchar(255)" json:"tahun"`
	Kondisi  string    `gorm:"column:kondisi;type:varchar(255)" json:"kondisi"`
	Latittude  string    `gorm:"column:latittude;type:varchar(255)" json:"latittude"`
	Longitude  string    `gorm:"column:longitude;type:varchar(255)" json:"longitude"`
	JumlahMesin  string    `gorm:"column:jumlah_mesin;type:varchar(255)" json:"jumlah_mesin"`
	CreatedAt   time.Time    `gorm:"column:created_at" json:"created_at"`	
	UpdatedAt   time.Time    `gorm:"column:updated_at" json:"updated_at"`
}

type DistributionSum struct {
	Daya_pasang 		int32 `gorm:"column:daya_pasang;type:integer" json:"daya_pasang"`
	Daya_mampu 		int32 `gorm:"column:daya_mampu;type:integer" json:"daya_mampu"`
	Sentral 	int32	 `json:"sentral"` 
	Mesin 	int32	 `json:"mesin"` 
}

// TableName return table name of Distribution model
func (Distribution) TableName() string {
	return "pltd_sebarans"
}

func (DistributionSum) TableName() string {
	return "pltd_sebarans"
}

