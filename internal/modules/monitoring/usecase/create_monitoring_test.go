package usecase

import (
	"context"

	"PLTD3/internal/modules/monitoring/domain"
	mockrepo "PLTD3/pkg/mocks/modules/monitoring/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_monitoringUsecaseImpl_CreateMonitoring(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		monitoringRepo := &mockrepo.MonitoringRepository{}
		monitoringRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MonitoringRepo").Return(monitoringRepo)

		uc := monitoringUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.CreateMonitoring(context.Background(), &domain.RequestMonitoring{})
		assert.NoError(t, err)
	})
}
