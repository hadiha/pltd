package usecase

import (
	"context"
	"fmt"

	"PLTD3/internal/modules/business/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *businessUsecaseImpl) GetConditionUnit(ctx context.Context, filter *domain.FilterBusiness) (result []domain.ResponseBusinessCondition, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BusinessUsecase:GetConditionUnit")
	defer trace.Finish()

	var data []shareddomain.Business
	data, err = uc.repoSQL.BusinessRepo().FetchConditionUnit(ctx, filter)
	if err != nil {
		return result, err
	}
	fmt.Println("condi", data)

	for _, detail := range data {
		result = append(result, domain.ResponseBusinessCondition{
			Status : detail.Status,
			Jumlah : detail.Jumlah,
		})
	}

	return
}
