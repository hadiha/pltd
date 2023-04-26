package usecase

import (
	"context"
	"time"

	"PLTD3/internal/modules/business/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *businessUsecaseImpl) GetDetailBusiness(ctx context.Context, id string) (result domain.ResponseBusiness, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BusinessUsecase:GetDetailBusiness")
	defer trace.Finish()

	var data shareddomain.Business
	repoFilter := domain.FilterBusiness{ID: &id}
	data, err = uc.repoSQL.BusinessRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.ID = data.ID
	// result.Field = data.Field
	result.CreatedAt = data.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = data.UpdatedAt.Format(time.RFC3339)
	return
}
