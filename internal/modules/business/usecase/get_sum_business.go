package usecase

import (
	"PLTD3/internal/modules/business/domain"
	"context"
	"fmt"

	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *businessUsecaseImpl) GetSumBusiness(ctx context.Context, filter *domain.FilterBusiness) (result domain.ResponseSumBusiness, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BusinessUsecase:GetSumBusiness")
	defer trace.Finish()

	var data shareddomain.Business
	data, err = uc.repoSQL.BusinessRepo().SumAllBusiness(ctx, filter)
	if err != nil {
		return result, err
	}
	fmt.Println("hai",data)

	result.Date = data.Date
	result.Dmn = data.Dmn
	result.Cad = data.Cad
	result.Dmp = data.Dmp
	result.Bp = data.Bp
	result.Cad_mw = data.Cad_mw
	result.Po = data.Po
	result.Mo = data.Mo
	result.Fd = data.Fd
	result.Total = data.Total

	return
}
