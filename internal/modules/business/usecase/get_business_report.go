package usecase

import (
	"context"
	"time"

	"PLTD3/internal/modules/business/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *businessUsecaseImpl) GetBusinessReport(ctx context.Context, filter *domain.FilterBusiness) (results []domain.ResponseBusinessReport, meta candishared.Meta, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BusinessUsecase:GetBusinessReport")
	defer trace.Finish()

	var data []shareddomain.BusinessReport
	data, err = uc.repoSQL.BusinessRepo().FetchAllReport(ctx, filter)
	if err != nil {
		return results, meta, err
	}
	count := uc.repoSQL.BusinessRepo().Count(ctx, filter)
	meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	for _, detail := range data {
		results = append(results, domain.ResponseBusinessReport{
			ID: detail.ID,
			Unit: detail.Unit,
			Day: detail.Day,
			Date: detail.CreatedAt.Format(time.RFC3339),
			Dmn: detail.Dmn,
			Cad: detail.Cad,
			Dmp: detail.Dmp,
			Bp: detail.Bp,
			Cad_mw: detail.Cad_mw,
			Status: detail.Status,
			Description: detail.Description,
			Po: detail.Po,
			Mo: detail.Mo,
			Fd: detail.Fd,
			Total: detail.Total,
			CreatedAt: detail.CreatedAt.Format(time.RFC3339),
			UpdatedAt: detail.UpdatedAt.Format(time.RFC3339),
		})
	}

	return
}
