// Code generated by candi v1.14.3.

package repository

import (
	"context"
	"database/sql"
	"fmt"

	// @candi:repositoryImport
	monitoringrepo "PLTD3/internal/modules/monitoring/repository"
	kpirepo "PLTD3/internal/modules/kpi/repository"
	distributionrepo "PLTD3/internal/modules/distribution/repository"
	consumptionrepo "PLTD3/internal/modules/consumption/repository"
	businessrepo "PLTD3/internal/modules/business/repository"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"

	"PLTD3/pkg/shared"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	// RepoSQL abstraction
	RepoSQL interface {
		WithTransaction(ctx context.Context, txFunc func(ctx context.Context) error) (err error)

		// @candi:repositoryMethod
		MonitoringRepo() monitoringrepo.MonitoringRepository
		KpiRepo() kpirepo.KpiRepository
		DistributionRepo() distributionrepo.DistributionRepository
		ConsumptionRepo() consumptionrepo.ConsumptionRepository
		BusinessRepo() businessrepo.BusinessRepository
	}

	repoSQLImpl struct {
		readDB, writeDB *gorm.DB
	
		// register all repository from modules
		// @candi:repositoryField
		monitoringRepo monitoringrepo.MonitoringRepository
		kpiRepo kpirepo.KpiRepository
		distributionRepo distributionrepo.DistributionRepository
		consumptionRepo consumptionrepo.ConsumptionRepository
		businessRepo businessrepo.BusinessRepository
	}
)

var (
	globalRepoSQL RepoSQL
)

// setSharedRepoSQL set the global singleton "RepoSQL" implementation
func setSharedRepoSQL(readDB, writeDB *sql.DB) {
	gormRead, err := gorm.Open(postgres.New(postgres.Config{
		Conn: readDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	gormWrite, err := gorm.Open(postgres.New(postgres.Config{
		Conn: writeDB,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}
	
	shared.AddGormCallbacks(gormRead)
	shared.AddGormCallbacks(gormWrite)

	globalRepoSQL = NewRepositorySQL(gormRead, gormWrite)
}

// GetSharedRepoSQL returns the global singleton "RepoSQL" implementation
func GetSharedRepoSQL() RepoSQL {
	return globalRepoSQL
}

// NewRepositorySQL constructor
func NewRepositorySQL(readDB, writeDB *gorm.DB) RepoSQL {

	return &repoSQLImpl{
		readDB: readDB, writeDB: writeDB,

		// @candi:repositoryConstructor
		monitoringRepo: monitoringrepo.NewMonitoringRepoSQL(readDB, writeDB),
		kpiRepo: kpirepo.NewKpiRepoSQL(readDB, writeDB),
		distributionRepo: distributionrepo.NewDistributionRepoSQL(readDB, writeDB),
		consumptionRepo: consumptionrepo.NewConsumptionRepoSQL(readDB, writeDB),
		businessRepo: businessrepo.NewBusinessRepoSQL(readDB, writeDB),
	}
}

// WithTransaction run transaction for each repository with context, include handle canceled or timeout context
func (r *repoSQLImpl) WithTransaction(ctx context.Context, txFunc func(ctx context.Context) error) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "RepoSQL:Transaction")
	defer trace.Finish()

	tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB)
	if !ok {
		tx = r.writeDB.Begin()
		if tx.Error != nil {
			return tx.Error
		}
		ctx = candishared.SetToContext(ctx, candishared.ContextKeySQLTransaction, tx)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			trace.SetError(err)
		} else {
			tx.Commit()
		}
	}()

	errChan := make(chan error)
	go func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: %v", r)
			}
			close(errChan)
		}()

		if err := txFunc(ctx); err != nil {
			errChan <- err
		}
	}(ctx)

	select {
	case <-ctx.Done():
		return fmt.Errorf("Canceled or timeout: %v", ctx.Err())
	case e := <-errChan:
		return e
	}
}

// @candi:repositoryImplementation
func (r *repoSQLImpl) MonitoringRepo() monitoringrepo.MonitoringRepository {
	return r.monitoringRepo
}

func (r *repoSQLImpl) KpiRepo() kpirepo.KpiRepository {
	return r.kpiRepo
}

func (r *repoSQLImpl) DistributionRepo() distributionrepo.DistributionRepository {
	return r.distributionRepo
}

func (r *repoSQLImpl) ConsumptionRepo() consumptionrepo.ConsumptionRepository {
	return r.consumptionRepo
}

func (r *repoSQLImpl) BusinessRepo() businessrepo.BusinessRepository {
	return r.businessRepo
}
