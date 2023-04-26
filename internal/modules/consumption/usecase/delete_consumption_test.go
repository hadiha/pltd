package usecase

import (
	"context"

	mockrepo "PLTD3/pkg/mocks/modules/consumption/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_consumptionUsecaseImpl_DeleteConsumption(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		consumptionRepo := &mockrepo.ConsumptionRepository{}
		consumptionRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ConsumptionRepo").Return(consumptionRepo)

		uc := consumptionUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteConsumption(context.Background(), "id")
		assert.NoError(t, err)
	})
}
