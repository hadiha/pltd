package usecase

import (
	"context"
	"time"

	"PLTD3/internal/modules/kpi/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *kpiUsecaseImpl) GetDetailKpi(ctx context.Context, id string) (result domain.ResponseKpi, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "KpiUsecase:GetDetailKpi")
	defer trace.Finish()

	var data shareddomain.Kpi
	repoFilter := domain.FilterKpi{ID: &id}
	data, err = uc.repoSQL.KpiRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.ID = data.ID
	result.Field = data.Field
	result.CreatedAt = data.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = data.UpdatedAt.Format(time.RFC3339)
	return
}
