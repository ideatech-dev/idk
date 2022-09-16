package recap

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ISheet interface {
	SourceGorm(ctx context.Context, db *gorm.DB) error
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

func (s *Sheet) SourceGorm(ctx context.Context, db *gorm.DB) error {
	tableName := db.Statement.Table
	dbName := db.Statement.Schema.Name

	columns, err := s.fetchColumnsMysql(ctx, db, tableName, dbName)
	if err != nil {
		return errors.Wrapf(err, "error fetch mysql columns")
	}
	s.columns = columns

	data, err := s.fetchData(ctx, db)
	if err != nil {
		return errors.Wrapf(err, "error fetch mysql data")
	}
	s.data = data

	return nil
}

func (s *Sheet) fetchColumnsMysql(ctx context.Context, db *gorm.DB, tableName, dbName string) (cols []string, err error) {
	query := db.WithContext(ctx).
		Select("COLUMN_NAME").
		Table("information_schema.columns").
		Where("table_name = ?", tableName).
		Where("table_schema = ?", dbName).
		Order("ORDINAL_POSITION")

	if err = query.Find(&cols).Error; err != nil {
		return
	}

	return
}

func (s *Sheet) fetchData(ctx context.Context, db *gorm.DB) (data []map[string]interface{}, err error) {
	err = db.WithContext(ctx).Find(&data).Error
	return
}
