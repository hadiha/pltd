package usecase

import (
	"context"

	"PLTD3/internal/modules/kpi/domain"
	mockrepo "PLTD3/pkg/mocks/modules/kpi/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_kpiUsecaseImpl_CreateKpi(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		kpiRepo := &mockrepo.KpiRepository{}
		kpiRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("KpiRepo").Return(kpiRepo)

		uc := kpiUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.CreateKpi(context.Background(), &domain.RequestKpi{})
		assert.NoError(t, err)
	})
}
