package usecase

import (
	"context"
	"fmt"

	"PLTD3/internal/modules/distribution/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *distributionUsecaseImpl) GetSumDistribution(ctx context.Context, filter *domain.FilterDistribution) (result domain.ResponseSumDistribution, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionUsecase:GetSumDistribution")
	defer trace.Finish()

	var data shareddomain.DistributionSum
	data, err = uc.repoSQL.DistributionRepo().SumAllDistribution(ctx, filter)
	if err != nil {
		return result, err
	}
	
	fmt.Println("ini",data)

	result.Daya_pasang = float32(data.Daya_pasang)
	result.Daya_mampu = float32(data.Daya_mampu)
	result.Sentral = float32(data.Sentral)
	result.Mesin = float32(data.Mesin)
	

	return
}
