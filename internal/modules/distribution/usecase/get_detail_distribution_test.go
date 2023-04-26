package usecase

import (
	"context"

	mockrepo "PLTD3/pkg/mocks/modules/distribution/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_distributionUsecaseImpl_GetDetailDistribution(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		distributionRepo := &mockrepo.DistributionRepository{}
		distributionRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Distribution{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("DistributionRepo").Return(distributionRepo)

		uc := distributionUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailDistribution(context.Background(), "id")
		assert.NoError(t, err)
	})
}
