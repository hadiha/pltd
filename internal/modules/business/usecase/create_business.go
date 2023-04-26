package usecase

import (
	"context"

	"PLTD3/internal/modules/business/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *businessUsecaseImpl) CreateBusiness(ctx context.Context, req *domain.RequestBusiness) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BusinessUsecase:CreateBusiness")
	defer trace.Finish()

	return uc.repoSQL.BusinessRepo().Save(ctx, &shareddomain.Business{
		// Field: req.Field,
		ID: req.ID,
	})
}
