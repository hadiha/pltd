// Code generated by candi v1.14.3.

package repository

import (
	"context"

	"PLTD3/internal/modules/monitoring/domain"
	shareddomain "PLTD3/pkg/shared/domain"
)

// MonitoringRepository abstract interface
type MonitoringRepository interface {
	SumAllMonitoring(ctx context.Context, filter *domain.FilterMonitoring) (shareddomain.Monitoring, error)

	FetchAll(ctx context.Context, filter *domain.FilterMonitoring) ([]shareddomain.Monitoring, error)
	Count(ctx context.Context, filter *domain.FilterMonitoring) int
	Find(ctx context.Context, filter *domain.FilterMonitoring) (shareddomain.Monitoring, error)
	Save(ctx context.Context, data *shareddomain.Monitoring) error
	Delete(ctx context.Context, data *shareddomain.Monitoring) (err error)
}
