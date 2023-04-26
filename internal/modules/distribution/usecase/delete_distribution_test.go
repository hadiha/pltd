package usecase

import (
	"context"

	mockrepo "PLTD3/pkg/mocks/modules/distribution/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_distributionUsecaseImpl_DeleteDistribution(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		distributionRepo := &mockrepo.DistributionRepository{}
		distributionRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("DistributionRepo").Return(distributionRepo)

		uc := distributionUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteDistribution(context.Background(), "id")
		assert.NoError(t, err)
	})
}
