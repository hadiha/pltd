package usecase

import (
	"context"

	"PLTD3/internal/modules/business/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *businessUsecaseImpl) UpdateBusiness(ctx context.Context, data *domain.RequestBusiness) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BusinessUsecase:UpdateBusiness")
	defer trace.Finish()

	repoFilter := domain.FilterBusiness{ID: &data.ID}
	existing, err := uc.repoSQL.BusinessRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	// existing.Field = data.Field
	err = uc.repoSQL.BusinessRepo().Save(ctx, &existing)
	return
}
