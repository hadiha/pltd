package usecase

import (
	"context"
	
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *consumptionUsecaseImpl) DeleteConsumption(ctx context.Context, id string) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ConsumptionUsecase:DeleteConsumption")
	defer trace.Finish()

	
	return uc.repoSQL.ConsumptionRepo().Delete(ctx, &shareddomain.Consumption{
		ID: id,
	})
}
