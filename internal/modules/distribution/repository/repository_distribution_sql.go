// Code generated by candi v1.14.3.

package repository

import (
	"context"
	"fmt"

	"time"

	"github.com/google/uuid"

	"PLTD3/internal/modules/distribution/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"

	"PLTD3/pkg/shared"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type distributionRepoSQL struct {
	readDB, writeDB *gorm.DB
}

// NewDistributionRepoSQL mongo repo constructor
func NewDistributionRepoSQL(readDB, writeDB *gorm.DB) DistributionRepository {
	return &distributionRepoSQL{
		readDB, writeDB,
	}
}

func (r *distributionRepoSQL) FetchAllByInstansi(ctx context.Context, filter *domain.FilterDistribution) (data []shareddomain.Distribution, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionRepoSQL:FetchAllByInstansi")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	if filter.OrderBy == "" {
		filter.OrderBy = "updated_at"
	}
	
	fmt.Println("tiga", &data)
	err = r.setFilterDistribution(shared.SetSpanToGorm(ctx, r.readDB), filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&data).Error
	return
}

func (r *distributionRepoSQL) FetchAll(ctx context.Context, filter *domain.FilterDistribution) (data []shareddomain.Distribution, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionRepoSQL:FetchAll")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	if filter.OrderBy == "" {
		filter.OrderBy = "updated_at"
	}
	
	err = r.setFilterDistribution(shared.SetSpanToGorm(ctx, r.readDB), filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&data).Error
	return
}

func (r *distributionRepoSQL) Count(ctx context.Context, filter *domain.FilterDistribution) (count int) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionRepoSQL:Count")
	defer trace.Finish()

	var total int64
	r.setFilterDistribution(shared.SetSpanToGorm(ctx, r.readDB), filter).Model(&shareddomain.Distribution{}).Count(&total)
	count = int(total)
	
	trace.Log("count", count)
	return
}

func (r *distributionRepoSQL) Find(ctx context.Context, filter *domain.FilterDistribution) (result shareddomain.Distribution, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionRepoSQL:Find")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	err = r.setFilterDistribution(shared.SetSpanToGorm(ctx, r.readDB), filter).First(&result).Error
	return
}

func (r *distributionRepoSQL) Save(ctx context.Context, data *shareddomain.Distribution) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionRepoSQL:Save")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()
	trace.Log("data", data)

	db := r.writeDB
	if tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB); ok {
		db = tx
	}
	data.UpdatedAt = time.Now()
	if data.CreatedAt.IsZero() {
		data.CreatedAt = time.Now()
	}
	if data.ID == "" {
		data.ID = uuid.NewString()
		err = shared.SetSpanToGorm(ctx, db).Create(data).Error
	} else {
		err = shared.SetSpanToGorm(ctx, db).Save(data).Error
	}
	return
}

func (r *distributionRepoSQL) Delete(ctx context.Context, data *shareddomain.Distribution) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DistributionRepoSQL:Delete")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	db := r.writeDB
	if tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB); ok {
		db = tx
	}
	err = shared.SetSpanToGorm(ctx, db).Delete(data).Error
	return
}

func (r *distributionRepoSQL) setFilterDistribution(db *gorm.DB, filter *domain.FilterDistribution) *gorm.DB {

	if filter.ID != nil {
		db = db.Where("id = ?", *filter.ID)
	}
	if filter.Instansi != nil {
		db = db.Where("instansi = ?", *filter.Instansi)
	}
	if filter.Search != "" {
		db = db.Where("(pembangkit ILIKE '%%' || ? || '%%')", filter.Search)
	}

	for _, preload := range filter.Preloads {
		db = db.Preload(preload)
	}

	return db
}
