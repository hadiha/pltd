package usecase

import (
	"context"
	"errors"

	"PLTD3/internal/modules/business/domain"
	mockrepo "PLTD3/pkg/mocks/modules/business/repository"
	mocksharedrepo "PLTD3/pkg/mocks/shared/repository"
	shareddomain "PLTD3/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_businessUsecaseImpl_UpdateBusiness(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		businessRepo := &mockrepo.BusinessRepository{}
		businessRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Business{}, nil)
		businessRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BusinessRepo").Return(businessRepo)

		uc := businessUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateBusiness(context.Background(), &domain.RequestBusiness{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		businessRepo := &mockrepo.BusinessRepository{}
		businessRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Business{}, errors.New("Error"))
		businessRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BusinessRepo").Return(businessRepo)

		uc := businessUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateBusiness(context.Background(), &domain.RequestBusiness{})
		assert.Error(t, err)
	})
}
