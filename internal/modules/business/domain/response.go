package domain

// ResponseBusiness model
type ResponseBusinessReport struct {
	ID        string `json:"id"`
	Unit 			string `json:"unit"`
	Day 			string `json:"day"`
	Date 			string `json:"date"`
	Dmn 			float32 `json:"dmn"`
	Cad 			float32 `json:"cad"`
	Dmp 			float32 `json:"dmp"`
	Bp 				float32 `json:"bp"`
	Cad_mw 		float32 `json:"cad_mw"`
	Status		string `json:"status"`
	Description	string `json:"description"`
	Po				float32 `json:"po"`
	Mo				float32 `json:"mo"`
	Fd				float32 `json:"fd"`
	Total			float32 `json:"total"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ResponseBusiness struct {
	ID  string `json:"iD"`
	Tgl_input  string `json:"tgl_input"`
	Tgl_laporan  string `json:"tgl_laporan"`
	Jenis_pembangkit  string `json:"jenis_pembangkit"`
	Shift  string `json:"shift"`
	Sh_ipp_ap  string `json:"sh_ipp_ap"`
	Sistem  string `json:"sistem"`
	Isolated  string `json:"isolated"`
	Sentral  string `json:"sentral"`
	Unit_mesin  string `json:"unit_mesin"`
	Tegangan  string `json:"tegangan"`
	Daya_pasang  string `json:"daya_pasang"`
	Daya_mampu  string `json:"daya_mampu"`
	Dmp  string `json:"dmp"`
	Beban  string `json:"beban"`
	Status  string `json:"status"`
	Deskripsi  string `json:"deskripsi"`
	Tgl_mulai  string `json:"tgl_mulai"`
	Tgl_selesai  string `json:"tgl_selesai"`
	Progress  string `json:"progress"`
	Outage  string `json:"outage"`
	Derating  string `json:"derating"`
	Hop_bb  string `json:"hop_bb"`
	Vol_bb  string `json:"vol_bb"`
	Vol_bbm  string `json:"vol_bbm"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ResponseBusinessCondition struct {
	Status string `json:"status"`
	Jumlah string `json:"jumlah"`
}