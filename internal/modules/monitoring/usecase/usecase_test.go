// Code generated by candi v1.14.3.

package usecase

import (
	"testing"

	mockdeps "github.com/golangid/candi/mocks/codebase/factory/dependency"
	mockinterfaces "github.com/golangid/candi/mocks/codebase/interfaces"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewMonitoringUsecase(t *testing.T) {
	mockPublisher := &mockinterfaces.Publisher{}
	mockBroker := &mockinterfaces.Broker{}
	mockBroker.On("GetPublisher").Return(mockPublisher)

	mockCache := &mockinterfaces.Cache{}
	mockRedisPool := &mockinterfaces.RedisPool{}
	mockRedisPool.On("Cache").Return(mockCache)

	mockDeps := &mockdeps.Dependency{}
	mockDeps.On("GetRedisPool").Return(mockRedisPool)
	mockDeps.On("GetBroker", mock.Anything).Return(mockBroker)

	uc, setFunc := NewMonitoringUsecase(mockDeps)
	setFunc(nil)
	assert.NotNil(t, uc)
}
