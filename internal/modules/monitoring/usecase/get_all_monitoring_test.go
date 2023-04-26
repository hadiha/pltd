package usecase

import (
	"context"
	"errors"

	"PLTD3/internal/modules/monitoring/domain"
	mockrepo "PLTD3/pkg/mocks/modules/monitoring/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_monitoringUsecaseImpl_GetAllMonitoring(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		monitoringRepo := &mockrepo.MonitoringRepository{}
		monitoringRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Monitoring{}, nil)
		monitoringRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MonitoringRepo").Return(monitoringRepo)

		uc := monitoringUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllMonitoring(context.Background(), &domain.FilterMonitoring{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		monitoringRepo := &mockrepo.MonitoringRepository{}
		monitoringRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Monitoring{}, errors.New("Error"))
		monitoringRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MonitoringRepo").Return(monitoringRepo)

		uc := monitoringUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllMonitoring(context.Background(), &domain.FilterMonitoring{})
		assert.Error(t, err)
	})
}
