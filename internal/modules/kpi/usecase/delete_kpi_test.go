package usecase

import (
	"context"

	mockrepo "PLTD3/pkg/mocks/modules/kpi/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_kpiUsecaseImpl_DeleteKpi(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		kpiRepo := &mockrepo.KpiRepository{}
		kpiRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("KpiRepo").Return(kpiRepo)

		uc := kpiUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteKpi(context.Background(), "id")
		assert.NoError(t, err)
	})
}
