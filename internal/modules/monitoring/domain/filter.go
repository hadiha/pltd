package domain

import "github.com/golangid/candi/candishared"

// FilterMonitoring model
type FilterMonitoring struct {
	candishared.Filter
	ID        *string `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Preloads  []string `json:"-"`
}
