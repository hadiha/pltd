package usecase

import (
	"context"
	"errors"

	"PLTD3/internal/modules/kpi/domain"
	mockrepo "PLTD3/pkg/mocks/modules/kpi/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_kpiUsecaseImpl_GetAllKpi(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		kpiRepo := &mockrepo.KpiRepository{}
		kpiRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Kpi{}, nil)
		kpiRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("KpiRepo").Return(kpiRepo)

		uc := kpiUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllKpi(context.Background(), &domain.FilterKpi{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		kpiRepo := &mockrepo.KpiRepository{}
		kpiRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Kpi{}, errors.New("Error"))
		kpiRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("KpiRepo").Return(kpiRepo)

		uc := kpiUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllKpi(context.Background(), &domain.FilterKpi{})
		assert.Error(t, err)
	})
}
