// Code generated by candi v1.14.3.

package usecase

import (
	"context"

	"PLTD3/internal/modules/consumption/domain"
	"PLTD3/pkg/shared/repository"
	"PLTD3/pkg/shared/usecase/common"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/factory/types"
	"github.com/golangid/candi/codebase/interfaces"
)

// ConsumptionUsecase abstraction
type ConsumptionUsecase interface {
	GetAllConsumption(ctx context.Context, filter *domain.FilterConsumption) (data []domain.ResponseConsumption, meta candishared.Meta, err error)
	GetDetailConsumption(ctx context.Context, id string) (data domain.ResponseConsumption, err error)
	CreateConsumption(ctx context.Context, data *domain.RequestConsumption) (err error)
	UpdateConsumption(ctx context.Context, data *domain.RequestConsumption) (err error)
	DeleteConsumption(ctx context.Context, id string) (err error)
}

type consumptionUsecaseImpl struct {
	sharedUsecase common.Usecase
	cache         interfaces.Cache
	repoSQL       repository.RepoSQL
	// repoMongo     repository.RepoMongo
	kafkaPub       interfaces.Publisher
}

// NewConsumptionUsecase usecase impl constructor
func NewConsumptionUsecase(deps dependency.Dependency) (ConsumptionUsecase, func(sharedUsecase common.Usecase)) {
	uc := &consumptionUsecaseImpl{
		repoSQL:   repository.GetSharedRepoSQL(),
		// repoMongo: repository.GetSharedRepoMongo(),
	}
	if redisPool := deps.GetRedisPool(); redisPool != nil {
		uc.cache = redisPool.Cache()
	}
	if kafkaBroker := deps.GetBroker(types.Kafka); kafkaBroker != nil {
		uc.kafkaPub = kafkaBroker.GetPublisher()
	}
	return uc, func(sharedUsecase common.Usecase) {
		uc.sharedUsecase = sharedUsecase
	}
}
