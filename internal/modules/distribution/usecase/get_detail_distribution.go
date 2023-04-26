package usecase

import (
	"context"
	"time"

	"PLTD3/internal/modules/distribution/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *distributionUsecaseImpl) GetDetailDistribution(ctx context.Context, id string) (result domain.ResponseDistribution, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionUsecase:GetDetailDistribution")
	defer trace.Finish()

	var data shareddomain.Distribution
	repoFilter := domain.FilterDistribution{ID: &id}
	data, err = uc.repoSQL.DistributionRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.ID = data.ID
	// result.Field = data.Field
	result.CreatedAt = data.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = data.UpdatedAt.Format(time.RFC3339)
	return
}
