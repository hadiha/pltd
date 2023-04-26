package usecase

import (
	"context"

	"PLTD3/internal/modules/kpi/domain"
	"github.com/golangid/candi/tracer"
)

func (uc *kpiUsecaseImpl) UpdateKpi(ctx context.Context, data *domain.RequestKpi) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "KpiUsecase:UpdateKpi")
	defer trace.Finish()

	repoFilter := domain.FilterKpi{ID: &data.ID}
	existing, err := uc.repoSQL.KpiRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Field = data.Field
	err = uc.repoSQL.KpiRepo().Save(ctx, &existing)
	return
}
