package usecase

import (
	"context"

	"PLTD3/internal/modules/business/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *businessUsecaseImpl) GetAllBusiness(ctx context.Context, filter *domain.FilterBusiness) (results []domain.ResponseBusiness, meta candishared.Meta, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BusinessUsecase:GetAllBusiness")
	defer trace.Finish()

	var data []shareddomain.Business
	data, err = uc.repoSQL.BusinessRepo().FetchAll(ctx, filter)
	if err != nil {
		return results, meta, err
	}
	count := uc.repoSQL.BusinessRepo().Count(ctx, filter)
	meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	for _, detail := range data {
		results = append(results, domain.ResponseBusiness{
			ID : detail.ID,
			Tgl_input : detail.Tgl_input,
			Tgl_laporan : detail.Tgl_laporan,
			Jenis_pembangkit : detail.Jenis_pembangkit,
			Shift : detail.Shift,
			Sh_ipp_ap : detail.Sh_ipp_ap,
			Sistem : detail.Sistem,
			Isolated : detail.Isolated,
			Sentral : detail.Sentral,
			Unit_mesin : detail.Unit_mesin,
			Tegangan : detail.Tegangan,
			Daya_pasang : detail.Daya_pasang,
			Daya_mampu : detail.Daya_mampu,
			Dmp : detail.Dmp,
			Beban : detail.Beban,
			Status : detail.Status,
			Deskripsi : detail.Deskripsi,
			Tgl_mulai : detail.Tgl_mulai,
			Tgl_selesai : detail.Tgl_selesai,
			Progress : detail.Progress,
			Outage : detail.Outage,
			Derating : detail.Derating,
			Hop_bb : detail.Hop_bb,
			Vol_bb : detail.Vol_bb,
			Vol_bbm : detail.Vol_bbm,
			UpdatedAt : detail.UpdatedAt.String(),
		})
	}

	return
}
