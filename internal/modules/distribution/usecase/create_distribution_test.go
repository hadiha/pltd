package usecase

import (
	"context"

	"PLTD3/internal/modules/distribution/domain"
	mockrepo "PLTD3/pkg/mocks/modules/distribution/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_distributionUsecaseImpl_CreateDistribution(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		distributionRepo := &mockrepo.DistributionRepository{}
		distributionRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("DistributionRepo").Return(distributionRepo)

		uc := distributionUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.CreateDistribution(context.Background(), &domain.RequestDistribution{})
		assert.NoError(t, err)
	})
}
