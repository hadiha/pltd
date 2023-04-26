package usecase

import (
	"PLTD3/internal/modules/monitoring/domain"
	"context"

	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *monitoringUsecaseImpl) GetSumMonitoring(ctx context.Context, filter *domain.FilterMonitoring) (result domain.ResponseSumMonitoring, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringUsecase:GetSumMonitoring")
	defer trace.Finish()

	var data shareddomain.Monitoring
	data, err = uc.repoSQL.MonitoringRepo().SumAllMonitoring(ctx, filter)
	if err != nil {
		return result, err
	}

	result.Date = data.Date
	result.Dmn = data.Dmn
	result.Cad = data.Cad
	result.Dmp = data.Dmp
	result.Bp = data.Bp
	result.Cad_mw = data.Cad_mw
	result.Po = data.Po
	result.Mo = data.Mo
	result.Fd = data.Fd
	result.Total = data.Total

	return
}
