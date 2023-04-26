package domain

import "github.com/golangid/candi/candishared"

// FilterConsumption model
type FilterConsumption struct {
	candishared.Filter
	ID        *string `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Preloads  []string `json:"-"`
}
