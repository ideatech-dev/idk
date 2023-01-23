package recap

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"testing"
)

func TestSheet_SourceGormPostgres(t *testing.T) {
	type fields struct {
		name    string
		columns []string
		data    []map[string]interface{}
	}
	type args struct {
		ctx         context.Context
		tableSchema string
		tableName   string
	}
	tests := []struct {
		name                 string
		isSkipping           bool
		fields               fields
		args                 args
		funcGenerateArgQuery func(t2 *testing.T) *gorm.DB
		wantErr              bool
	}{
		{
			name:       "Postgres - Talenthub",
			isSkipping: true,
			args: args{
				ctx:         context.Background(),
				tableSchema: "public",
				tableName:   "_download_programme_participant",
			},
			funcGenerateArgQuery: func(t *testing.T) *gorm.DB {
				dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
					"<HOST>",
					"<USERNAME>",
					"<PASS>",
					"<DB_NAME>",
					"<DB_PORT>",
				)
				db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
					Logger: logger.Default.LogMode(logger.Info),
				})
				if err != nil {
					t.Error(err)
				}

				db = db.Where("programme_id = ?", 59)

				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sheet{
				name:    tt.fields.name,
				columns: tt.fields.columns,
				data:    tt.fields.data,
			}

			if tt.isSkipping {
				t.Skip()
			}

			query := tt.funcGenerateArgQuery(t)

			if err := s.SourceGormPostgres(tt.args.ctx, tt.args.tableSchema, tt.args.tableName, query); (err != nil) != tt.wantErr {
				t.Errorf("SourceGorm() error = %v, wantErr %v", err, tt.wantErr)
			}

			log.Printf("columns: %v \n", s.columns)
			log.Printf("data: %v \n", s.data)
		})
	}
}
