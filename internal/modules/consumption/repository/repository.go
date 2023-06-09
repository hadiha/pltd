// Code generated by candi v1.14.3.

package repository

import (
	"context"

	"PLTD3/internal/modules/consumption/domain"
	shareddomain "PLTD3/pkg/shared/domain"
)

// ConsumptionRepository abstract interface
type ConsumptionRepository interface {
	FetchAll(ctx context.Context, filter *domain.FilterConsumption) ([]shareddomain.Consumption, error)
	Count(ctx context.Context, filter *domain.FilterConsumption) int
	Find(ctx context.Context, filter *domain.FilterConsumption) (shareddomain.Consumption, error)
	Save(ctx context.Context, data *shareddomain.Consumption) error
	Delete(ctx context.Context, data *shareddomain.Consumption) (err error)
}
