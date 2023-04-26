package usecase

import (
	"context"
	
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *businessUsecaseImpl) DeleteBusiness(ctx context.Context, id string) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BusinessUsecase:DeleteBusiness")
	defer trace.Finish()

	
	return uc.repoSQL.BusinessRepo().Delete(ctx, &shareddomain.Business{
		ID: id,
	})
}
