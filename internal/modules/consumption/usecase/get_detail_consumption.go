package usecase

import (
	"context"
	"time"

	"PLTD3/internal/modules/consumption/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *consumptionUsecaseImpl) GetDetailConsumption(ctx context.Context, id string) (result domain.ResponseConsumption, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ConsumptionUsecase:GetDetailConsumption")
	defer trace.Finish()

	var data shareddomain.Consumption
	repoFilter := domain.FilterConsumption{ID: &id}
	data, err = uc.repoSQL.ConsumptionRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.ID = data.ID
	// result.Field = data.Field
	result.CreatedAt = data.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = data.UpdatedAt.Format(time.RFC3339)
	return
}
