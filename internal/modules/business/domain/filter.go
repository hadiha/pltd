package domain

import "github.com/golangid/candi/candishared"

// FilterBusiness model
type FilterBusiness struct {
	candishared.Filter
	ID        *string `json:"id"`
	Unit        *string `json:"unit"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Preloads  []string `json:"-"`
}
