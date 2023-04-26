package usecase

import (
	"context"

	"PLTD3/internal/modules/distribution/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *distributionUsecaseImpl) CreateDistribution(ctx context.Context, req *domain.RequestDistribution) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionUsecase:CreateDistribution")
	defer trace.Finish()

	return uc.repoSQL.DistributionRepo().Save(ctx, &shareddomain.Distribution{
		// Field: req.Field,
	})
}
