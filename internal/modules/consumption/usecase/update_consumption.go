package usecase

import (
	"context"

	"PLTD3/internal/modules/consumption/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *consumptionUsecaseImpl) UpdateConsumption(ctx context.Context, data *domain.RequestConsumption) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ConsumptionUsecase:UpdateConsumption")
	defer trace.Finish()

	repoFilter := domain.FilterConsumption{ID: &data.ID}
	existing, err := uc.repoSQL.ConsumptionRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	// existing.Field = data.Field
	err = uc.repoSQL.ConsumptionRepo().Save(ctx, &existing)
	return
}
