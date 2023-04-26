package usecase

import (
	"context"

	mockrepo "PLTD3/pkg/mocks/modules/business/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_businessUsecaseImpl_GetDetailBusiness(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		businessRepo := &mockrepo.BusinessRepository{}
		businessRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Business{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BusinessRepo").Return(businessRepo)

		uc := businessUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailBusiness(context.Background(), "id")
		assert.NoError(t, err)
	})
}
