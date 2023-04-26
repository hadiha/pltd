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

func Test_businessUsecaseImpl_GetAllBusiness(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		businessRepo := &mockrepo.BusinessRepository{}
		businessRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Business{}, nil)
		businessRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BusinessRepo").Return(businessRepo)

		uc := businessUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllBusiness(context.Background(), &domain.FilterBusiness{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		businessRepo := &mockrepo.BusinessRepository{}
		businessRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Business{}, errors.New("Error"))
		businessRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BusinessRepo").Return(businessRepo)

		uc := businessUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllBusiness(context.Background(), &domain.FilterBusiness{})
		assert.Error(t, err)
	})
}
