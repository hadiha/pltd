package usecase

import (
	"context"
	"time"

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
			ID : detail.ID,
			NamaRegional : detail.NamaRegional,
			NomorRegional : detail.NomorRegional,
			NamaSistem : detail.NamaSistem,
			KodeSistem : detail.KodeSistem,
			NamaUnit : detail.NamaUnit,
			KodeUnit : detail.KodeUnit,
			NamaPembangkit : detail.NamaPembangkit,
			NomorPembangkit : detail.NomorPembangkit,
			Instansi : detail.Instansi,
			Provinsi : detail.Provinsi,
			Alamat : detail.Alamat,
			JenisKit : detail.JenisKit,
			Bbm : detail.Bbm,
			DayaPasang : detail.DayaPasang,
			DayaMampu : detail.DayaMampu,
			Tahun : detail.Tahun,
			Kondisi : detail.Kondisi,
			Latittude : detail.Latittude,
			Longitude : detail.Longitude,
			JumlahMesin : detail.JumlahMesin,
			CreatedAt : detail.CreatedAt.Format(time.RFC3339),
			UpdatedAt : detail.UpdatedAt.Format(time.RFC3339),
		})
	}

	return
}
