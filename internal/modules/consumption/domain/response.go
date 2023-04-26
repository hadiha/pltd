package domain

// ResponseConsumption model
type ResponseConsumption struct {
	ID       	 	string `json:"id"`
	Unit 				string `json:"unit"`
	Pembangkit 	string `json:"pembangkit"`
	Minyak 			int32 `json:"minyak"`
	Gas 				int32 `json:"gas"`
	Batubara 		int32 `json:"batubara"`
	Biomass 		int32 `json:"biomass"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	// Child []ResponseBbm `json:"child,omitempty"`
}
