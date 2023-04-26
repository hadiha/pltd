package usecase

import (
	"context"

	"PLTD3/internal/modules/monitoring/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *monitoringUsecaseImpl) UpdateMonitoring(ctx context.Context, data *domain.RequestMonitoring) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringUsecase:UpdateMonitoring")
	defer trace.Finish()

	repoFilter := domain.FilterMonitoring{ID: &data.ID}
	existing, err := uc.repoSQL.MonitoringRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	// existing.Field = data.Field
	err = uc.repoSQL.MonitoringRepo().Save(ctx, &existing)
	return
}
