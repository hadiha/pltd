package usecase

import (
	"context"
	"time"

	"PLTD3/internal/modules/kpi/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *kpiUsecaseImpl) GetAllKpi(ctx context.Context, filter *domain.FilterKpi) (results []domain.ResponseKpi, meta candishared.Meta, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "KpiUsecase:GetAllKpi")
	defer trace.Finish()

	var data []shareddomain.Kpi
	data, err = uc.repoSQL.KpiRepo().FetchAll(ctx, filter)
	if err != nil {
		return results, meta, err
	}
	count := uc.repoSQL.KpiRepo().Count(ctx, filter)
	meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	for _, detail := range data {
		results = append(results, domain.ResponseKpi{
			ID: detail.ID,
			Field: detail.Field,
			CreatedAt: detail.CreatedAt.Format(time.RFC3339),
			UpdatedAt: detail.UpdatedAt.Format(time.RFC3339),
		})
	}

	return
}
