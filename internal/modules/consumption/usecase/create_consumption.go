package usecase

import (
	"context"

	"PLTD3/internal/modules/consumption/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *consumptionUsecaseImpl) CreateConsumption(ctx context.Context, req *domain.RequestConsumption) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ConsumptionUsecase:CreateConsumption")
	defer trace.Finish()

	return uc.repoSQL.ConsumptionRepo().Save(ctx, &shareddomain.Consumption{
		// Field: req.Field,
	})
}
