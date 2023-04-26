package usecase

import (
	"context"

	mockrepo "PLTD3/pkg/mocks/modules/monitoring/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_monitoringUsecaseImpl_GetDetailMonitoring(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		monitoringRepo := &mockrepo.MonitoringRepository{}
		monitoringRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Monitoring{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MonitoringRepo").Return(monitoringRepo)

		uc := monitoringUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailMonitoring(context.Background(), "id")
		assert.NoError(t, err)
	})
}
