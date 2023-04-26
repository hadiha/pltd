package usecase

import (
	"context"
	"time"

	"PLTD3/internal/modules/monitoring/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *monitoringUsecaseImpl) GetAllMonitoring(ctx context.Context, filter *domain.FilterMonitoring) (results []domain.ResponseMonitoring, meta candishared.Meta, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringUsecase:GetAllMonitoring")
	defer trace.Finish()

	var data []shareddomain.Monitoring
	data, err = uc.repoSQL.MonitoringRepo().FetchAll(ctx, filter)
	if err != nil {
		return results, meta, err
	}
	count := uc.repoSQL.MonitoringRepo().Count(ctx, filter)
	meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	for _, detail := range data {
		results = append(results, domain.ResponseMonitoring{
			ID: detail.ID,
			// Field: detail.Field,
			CreatedAt: detail.CreatedAt.Format(time.RFC3339),
			UpdatedAt: detail.UpdatedAt.Format(time.RFC3339),
		})
	}

	return
}
