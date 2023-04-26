package usecase

import (
	"context"
	"errors"

	"PLTD3/internal/modules/distribution/domain"
	mockrepo "PLTD3/pkg/mocks/modules/distribution/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_distributionUsecaseImpl_UpdateDistribution(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		distributionRepo := &mockrepo.DistributionRepository{}
		distributionRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Distribution{}, nil)
		distributionRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("DistributionRepo").Return(distributionRepo)

		uc := distributionUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateDistribution(context.Background(), &domain.RequestDistribution{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		distributionRepo := &mockrepo.DistributionRepository{}
		distributionRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Distribution{}, errors.New("Error"))
		distributionRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("DistributionRepo").Return(distributionRepo)

		uc := distributionUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateDistribution(context.Background(), &domain.RequestDistribution{})
		assert.Error(t, err)
	})
}
