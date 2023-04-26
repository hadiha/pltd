package usecase

import (
	"context"
	"time"

	"PLTD3/internal/modules/monitoring/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *monitoringUsecaseImpl) GetDetailMonitoring(ctx context.Context, id string) (result domain.ResponseMonitoring, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringUsecase:GetDetailMonitoring")
	defer trace.Finish()

	var data shareddomain.Monitoring
	repoFilter := domain.FilterMonitoring{ID: &id}
	data, err = uc.repoSQL.MonitoringRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.ID = data.ID
	// result.Field = data.Field
	result.CreatedAt = data.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = data.UpdatedAt.Format(time.RFC3339)
	return
}
