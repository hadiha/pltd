package usecase

import (
	"context"

	mockrepo "PLTD3/pkg/mocks/modules/monitoring/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_monitoringUsecaseImpl_DeleteMonitoring(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		monitoringRepo := &mockrepo.MonitoringRepository{}
		monitoringRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MonitoringRepo").Return(monitoringRepo)

		uc := monitoringUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteMonitoring(context.Background(), "id")
		assert.NoError(t, err)
	})
}
