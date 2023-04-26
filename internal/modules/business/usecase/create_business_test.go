package usecase

import (
	"context"

	"PLTD3/internal/modules/business/domain"
	mockrepo "PLTD3/pkg/mocks/modules/business/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_businessUsecaseImpl_CreateBusiness(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		businessRepo := &mockrepo.BusinessRepository{}
		businessRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BusinessRepo").Return(businessRepo)

		uc := businessUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.CreateBusiness(context.Background(), &domain.RequestBusiness{})
		assert.NoError(t, err)
	})
}
