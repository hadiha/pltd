package usecase

import (
	"context"
	
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *kpiUsecaseImpl) DeleteKpi(ctx context.Context, id string) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "KpiUsecase:DeleteKpi")
	defer trace.Finish()

	
	return uc.repoSQL.KpiRepo().Delete(ctx, &shareddomain.Kpi{
		ID: id,
	})
}
