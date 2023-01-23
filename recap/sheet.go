package recap

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ISheet interface {
	SourceGormPostgres(ctx context.Context, tableSchema, tableName string, query *gorm.DB) error
	SourceGormMysql(ctx context.Context, dbName, tableName string, query *gorm.DB) error
}

var _ ISheet = &Sheet{}

func NewSheet(name string) *Sheet {
	return &Sheet{
		name: name,
	}
}

type Sheet struct {
	name    string
	columns []string
	data    []map[string]interface{}
}

func (s *Sheet) SourceGormPostgres(ctx context.Context, tableSchema, tableName string, query *gorm.DB) error {
	return s.sourceGorm(ctx, tableSchema, tableName, query)
}

func (s *Sheet) SourceGormMysql(ctx context.Context, dbName, tableName string, query *gorm.DB) error {
	return s.sourceGorm(ctx, dbName, tableName, query)
}

func (s *Sheet) sourceGorm(ctx context.Context, tableSchema, tableName string, query *gorm.DB) error {
	columns, err := s.fetchColumnsMysql(ctx, query, tableName, tableSchema)
	if err != nil {
		return errors.Wrapf(err, "error fetch mysql columns")
	}
	s.columns = columns

	data, err := s.fetchData(ctx, tableName, query)
	if err != nil {
		return errors.Wrapf(err, "error fetch mysql data")
	}
	s.data = data

	return nil
}

func (s *Sheet) fetchColumnsMysql(ctx context.Context, db *gorm.DB, tableName, tableSchema string) (cols []string, err error) {
	query := db.WithContext(ctx).Raw(`
		SELECT COLUMN_NAME
		FROM information_schema.columns
		WHERE table_name = ?
		AND table_schema = ?
		ORDER BY ORDINAL_POSITION
	`, tableName, tableSchema)

	if err = query.Scan(&cols).Error; err != nil {
		return
	}

	return
}

func (s *Sheet) fetchData(ctx context.Context, tableName string, db *gorm.DB) (data []map[string]interface{}, err error) {
	err = db.WithContext(ctx).Table(tableName).Find(&data).Error
	return
}
