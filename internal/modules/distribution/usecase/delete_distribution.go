package usecase

import (
	"context"
	
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *distributionUsecaseImpl) DeleteDistribution(ctx context.Context, id string) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionUsecase:DeleteDistribution")
	defer trace.Finish()

	
	return uc.repoSQL.DistributionRepo().Delete(ctx, &shareddomain.Distribution{
		ID: id,
	})
}
