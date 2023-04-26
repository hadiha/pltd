package domain

import "github.com/golangid/candi/candishared"

// FilterKpi model
type FilterKpi struct {
	candishared.Filter
	ID        *string `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Preloads  []string `json:"-"`
}
