package usecase

import (
	"context"

	mockrepo "PLTD3/pkg/mocks/modules/kpi/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_kpiUsecaseImpl_GetDetailKpi(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		kpiRepo := &mockrepo.KpiRepository{}
		kpiRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Kpi{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("KpiRepo").Return(kpiRepo)

		uc := kpiUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailKpi(context.Background(), "id")
		assert.NoError(t, err)
	})
}
