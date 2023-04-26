package usecase

import (
	"context"

	"PLTD3/internal/modules/kpi/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *kpiUsecaseImpl) CreateKpi(ctx context.Context, req *domain.RequestKpi) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "KpiUsecase:CreateKpi")
	defer trace.Finish()

	return uc.repoSQL.KpiRepo().Save(ctx, &shareddomain.Kpi{
		Field: req.Field,
	})
}
