package usecase

import (
	"context"

	"PLTD3/internal/modules/distribution/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *distributionUsecaseImpl) UpdateDistribution(ctx context.Context, data *domain.RequestDistribution) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionUsecase:UpdateDistribution")
	defer trace.Finish()

	repoFilter := domain.FilterDistribution{ID: &data.ID}
	existing, err := uc.repoSQL.DistributionRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	// existing.Field = data.Field
	err = uc.repoSQL.DistributionRepo().Save(ctx, &existing)
	return
}
