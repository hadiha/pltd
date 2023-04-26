package usecase

import (
	"context"

	mockrepo "PLTD3/pkg/mocks/modules/consumption/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_consumptionUsecaseImpl_GetDetailConsumption(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		consumptionRepo := &mockrepo.ConsumptionRepository{}
		consumptionRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Consumption{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ConsumptionRepo").Return(consumptionRepo)

		uc := consumptionUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailConsumption(context.Background(), "id")
		assert.NoError(t, err)
	})
}
