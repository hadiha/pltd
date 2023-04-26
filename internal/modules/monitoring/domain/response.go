package domain

// ResponseMonitoring model
type ResponseMonitoring struct {
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


type ResponseSumMonitoring struct {
	Date 			string `json:"date"`
	Dmn 			float32 `json:"dmn"`
	Cad 			float32 `json:"cad"`
	Dmp 			float32 `json:"dmp"`
	Bp 				float32 `json:"bp"`
	Cad_mw 		float32 `json:"cad_mw"`
	Po				float32 `json:"po"`
	Mo				float32 `json:"mo"`
	Fd				float32 `json:"fd"`
	Total			float32 `json:"total"`
}
