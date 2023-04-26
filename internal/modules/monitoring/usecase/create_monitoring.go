package usecase

import (
	"context"

	"PLTD3/internal/modules/monitoring/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *monitoringUsecaseImpl) CreateMonitoring(ctx context.Context, req *domain.RequestMonitoring) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringUsecase:CreateMonitoring")
	defer trace.Finish()

	return uc.repoSQL.MonitoringRepo().Save(ctx, &shareddomain.Monitoring{
		// Field: req.Field,
	})
}
