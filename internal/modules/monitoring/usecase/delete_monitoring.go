package usecase

import (
	"context"
	
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *monitoringUsecaseImpl) DeleteMonitoring(ctx context.Context, id string) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringUsecase:DeleteMonitoring")
	defer trace.Finish()

	
	return uc.repoSQL.MonitoringRepo().Delete(ctx, &shareddomain.Monitoring{
		ID: id,
	})
}
