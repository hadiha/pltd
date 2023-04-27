package usecase

import (
	"context"
	"fmt"

	"PLTD3/internal/modules/distribution/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *distributionUsecaseImpl) GetMapsDistribution(ctx context.Context, filter *domain.FilterDistribution) (result []domain.ResponseMapsDistribution, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionUsecase:GetSumDistribution")
	defer trace.Finish()

	var data []shareddomain.Distribution
	data, err = uc.repoSQL.DistributionRepo().MapsDistribution(ctx, filter)
	if err != nil {
		return result, err
	}
	
	fmt.Println("ini",data)

	for _, detail := range data {
		result = append(result, domain.ResponseMapsDistribution{
			NamaPembangkit : detail.NamaPembangkit,
			Instansi : detail.Instansi,
			Latittude : detail.Latittude,
			Longitude : detail.Longitude,
		})
	}

	return
}
