package domain

// ResponseDistribution model
type ResponseDistribution struct {
	ID       	 	string `json:"id"`
	Instansi 				string `json:"instansi"`
	Pembangkit 	string `json:"pembangkit"`
	Provinsi 			string `json:"provinsi"`
	Jenis_kit				string `json:"jenis_kit"`
	Bbm 		string `json:"bbm"`
	Daya_pasang 		int32 `json:"daya_pasang"`
	Daya_mampu 		int32 `json:"daya_mampu"`
	Tahun 		string `json:"tahun"`
	Kondisi 		string `json:"kondisi"`
	Lattitude 		string `json:"lattitude"`
	Longitude 		string `json:"longitude"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}


type ResponseSumDistribution struct {
	Daya_pasang 			float32 `json:"daya_pasang"`
	Daya_mampu 			float32 `json:"daya_mampu"`
	Sentral 			float32 `json:"sentral"`
	Mesin 				float32 `json:"mesin"`
}