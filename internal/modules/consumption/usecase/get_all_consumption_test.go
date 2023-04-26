package usecase

import (
	"context"
	"errors"

	"PLTD3/internal/modules/consumption/domain"
	mockrepo "PLTD3/pkg/mocks/modules/consumption/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_consumptionUsecaseImpl_GetAllConsumption(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		consumptionRepo := &mockrepo.ConsumptionRepository{}
		consumptionRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Consumption{}, nil)
		consumptionRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ConsumptionRepo").Return(consumptionRepo)

		uc := consumptionUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllConsumption(context.Background(), &domain.FilterConsumption{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		consumptionRepo := &mockrepo.ConsumptionRepository{}
		consumptionRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Consumption{}, errors.New("Error"))
		consumptionRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ConsumptionRepo").Return(consumptionRepo)

		uc := consumptionUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllConsumption(context.Background(), &domain.FilterConsumption{})
		assert.Error(t, err)
	})
}
