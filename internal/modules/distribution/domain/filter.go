package domain

import "github.com/golangid/candi/candishared"

// FilterDistribution model
type FilterDistribution struct {
	candishared.Filter
	ID        *string `json:"id"`
	Instansi 	*string `json:"instansi"`
	Pembangkit 	*string `json:"pembangkit"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Preloads  []string `json:"-"`
}
