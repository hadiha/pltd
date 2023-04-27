package domain

// ResponseDistribution model
type ResponseDistribution struct {
	ID  string `json:"id"`
	NamaRegional  string `json:"nama_regional"`
	NomorRegional  string `json:"nomor_regional"`
	NamaSistem  string `json:"nama_sistem"`
	KodeSistem  string `json:"kode_sistem"`
	NamaUnit  string `json:"nama_unit"`
	KodeUnit  string `json:"kode_unit"`
	NamaPembangkit  string `json:"nama_pembangkit"`
	NomorPembangkit  string `json:"nomor_pembangkit"`
	Instansi  string `json:"instansi"`
	Provinsi  string `json:"provinsi"`
	Alamat  string `json:"alamat"`
	JenisKit  string `json:"jenis_kit"`
	Bbm  string `json:"bbm"`
	DayaPasang  string `json:"daya_pasang"`
	DayaMampu  string `json:"daya_mampu"`
	Tahun  string `json:"tahun"`
	Kondisi  string `json:"kondisi"`
	Latittude  string `json:"latittude"`
	Longitude  string `json:"longitude"`
	JumlahMesin  string `json:"jumlah_mesin"`
	CreatedAt  string `json:"created_at"`	
	UpdatedAt  string `json:"updated_at"`
}


type ResponseSumDistribution struct {
	Daya_pasang 			float32 `json:"daya_pasang"`
	Daya_mampu 			float32 `json:"daya_mampu"`
	Sentral 			float32 `json:"sentral"`
	Mesin 				float32 `json:"mesin"`
}
type ResponseMapsDistribution struct {
	Instansi 			string `json:"instansi"`
	NamaPembangkit 			string `json:"nama_pembangkit"`
	Latittude 			string `json:"latittude"`
	Longitude 			string `json:"longitude"`
}