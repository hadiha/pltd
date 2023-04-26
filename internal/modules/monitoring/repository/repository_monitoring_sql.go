// Code generated by candi v1.14.3.

package repository

import (
	"context"

	"time"

	"github.com/google/uuid"

	"PLTD3/internal/modules/monitoring/domain"
	shareddomain "PLTD3/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"

	"PLTD3/pkg/shared"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type monitoringRepoSQL struct {
	readDB, writeDB *gorm.DB
}

// NewMonitoringRepoSQL mongo repo constructor
func NewMonitoringRepoSQL(readDB, writeDB *gorm.DB) MonitoringRepository {
	return &monitoringRepoSQL{
		readDB, writeDB,
	}
}

func (r *monitoringRepoSQL) FetchAll(ctx context.Context, filter *domain.FilterMonitoring) (data []shareddomain.Monitoring, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringRepoSQL:FetchAll")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	if filter.OrderBy == "" {
		filter.OrderBy = "updated_at"
	}
	
	err = r.setFilterMonitoring(shared.SetSpanToGorm(ctx, r.readDB), filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&data).Error
	return
}

func (r *monitoringRepoSQL) Count(ctx context.Context, filter *domain.FilterMonitoring) (count int) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringRepoSQL:Count")
	defer trace.Finish()

	var total int64
	r.setFilterMonitoring(shared.SetSpanToGorm(ctx, r.readDB), filter).Model(&shareddomain.Monitoring{}).Count(&total)
	count = int(total)
	
	trace.Log("count", count)
	return
}

func (r *monitoringRepoSQL) SumAllMonitoring(ctx context.Context, filter *domain.FilterMonitoring) (result shareddomain.Monitoring, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringRepoSQL:SumAllMonitoring")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	if filter.OrderBy == "" {
		filter.OrderBy = "updated_at"
	}
	

	// today := time.Now()
	// yesterday := today.AddDate(0, 0, -1)
	yesterday := "2023-04-10"
	err = r.setFilterMonitoring(shared.SetSpanToGorm(ctx, r.readDB), filter).Select(
		"date", 
		"SUM(dmn) as dmn", 
		"SUM(cad) as cad", 
		"SUM(dmp) as dmp", 
		"SUM(bp) as bp",
		"SUM(cad_mw) as cad_mw",
		"SUM(po) as po",
		"SUM(mo) as mo",
		"SUM(fd) as fd",
		"SUM(total) as total").Where("date = ?", yesterday).Group("date").Find(&result).Error
		// "SUM(total) as total").Where("date = ?", yesterday.Format("2006-01-02")).Group("date").Find(&result).Error
	return
}







func (r *monitoringRepoSQL) Find(ctx context.Context, filter *domain.FilterMonitoring) (result shareddomain.Monitoring, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringRepoSQL:Find")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	err = r.setFilterMonitoring(shared.SetSpanToGorm(ctx, r.readDB), filter).First(&result).Error
	return
}

func (r *monitoringRepoSQL) Save(ctx context.Context, data *shareddomain.Monitoring) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringRepoSQL:Save")
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

func (r *monitoringRepoSQL) Delete(ctx context.Context, data *shareddomain.Monitoring) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MonitoringRepoSQL:Delete")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	db := r.writeDB
	if tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB); ok {
		db = tx
	}
	err = shared.SetSpanToGorm(ctx, db).Delete(data).Error
	return
}

func (r *monitoringRepoSQL) setFilterMonitoring(db *gorm.DB, filter *domain.FilterMonitoring) *gorm.DB {

	if filter.ID != nil {
		db = db.Where("id = ?", *filter.ID)
	}
	if filter.Search != "" {
		db = db.Where("(field ILIKE '%%' || ? || '%%')", filter.Search)
	}

	for _, preload := range filter.Preloads {
		db = db.Preload(preload)
	}

	return db
}
