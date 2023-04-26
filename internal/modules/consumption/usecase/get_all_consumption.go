package usecase

import (
	"context"

	"PLTD3/internal/modules/consumption/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *consumptionUsecaseImpl) GetAllConsumption(ctx context.Context, filter *domain.FilterConsumption) (results []domain.ResponseConsumption, meta candishared.Meta, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ConsumptionUsecase:GetAllConsumption")
	defer trace.Finish()

	var data []shareddomain.Consumption
	data, err = uc.repoSQL.ConsumptionRepo().FetchAll(ctx, filter)
	if err != nil {
		return results, meta, err
	}
	count := uc.repoSQL.ConsumptionRepo().Count(ctx, filter)
	meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	for _, detail := range data {
		results = append(results, domain.ResponseConsumption{
			ID: detail.ID,
			Unit: detail.Unit,
			Pembangkit: detail.Pembangkit,
			Minyak: detail.Minyak,
			Gas: detail.Gas,
			Batubara: detail.Batubara,
			Biomass: detail.Biomass,
		})
	}

	return
}
