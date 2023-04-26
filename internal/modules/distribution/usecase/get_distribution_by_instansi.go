package usecase

import (
	"context"
	"fmt"

	"PLTD3/internal/modules/distribution/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *distributionUsecaseImpl) GetDistributionByInstansi(ctx context.Context, instansi string, filter *domain.FilterDistribution) (results []domain.ResponseDistribution, meta candishared.Meta, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionUsecase:GetDistributionByInstansi")
	defer trace.Finish()

	var data []shareddomain.Distribution
	repoFilter := domain.FilterDistribution{Instansi: &instansi}
	fmt.Println("dua", &instansi)
	data, err = uc.repoSQL.DistributionRepo().FetchAllByInstansi(ctx, &repoFilter)
	if err != nil {
		return results, meta, err
	}
	fmt.Println("empat", data)
	count := uc.repoSQL.DistributionRepo().Count(ctx, filter)
	meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	for _, detail := range data {
		results = append(results, domain.ResponseDistribution{
			ID: detail.ID,
			Instansi: detail.Instansi,
			Pembangkit: detail.Pembangkit,
			Provinsi: detail.Provinsi,
			Jenis_kit: detail.Jenis_kit,
			Bbm: detail.Bbm,
			Daya_pasang: detail.Daya_pasang,
			Daya_mampu: detail.Daya_mampu,
			Tahun: detail.Tahun,
			Kondisi: detail.Kondisi,
			Lattitude: detail.Lattitude,
			Longitude: detail.Longitude,
		})
	}

	return
}
