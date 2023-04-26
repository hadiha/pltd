package usecase

import (
	"context"

	"PLTD3/internal/modules/monitoring/domain"
	mockrepo "PLTD3/pkg/mocks/modules/monitoring/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_monitoringUsecaseImpl_GetSumMonitoring(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		monitoringRepo := &mockrepo.MonitoringRepository{}
		monitoringRepo.On("SumAllMonitoring", mock.Anything, mock.Anything).Return(shareddomain.Monitoring{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MonitoringRepo").Return(monitoringRepo)

		uc := monitoringUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetSumMonitoring(context.Background(), &domain.FilterMonitoring{})
		assert.NoError(t, err)
	})
}
