package usecase

import (
	"context"

	"PLTD3/internal/modules/distribution/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *distributionUsecaseImpl) GetAllDistribution(ctx context.Context, filter *domain.FilterDistribution) (results []domain.ResponseDistribution, meta candishared.Meta, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionUsecase:GetAllDistribution")
	defer trace.Finish()

	var data []shareddomain.Distribution
	data, err = uc.repoSQL.DistributionRepo().FetchAll(ctx, filter)
	if err != nil {
		return results, meta, err
	}
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
