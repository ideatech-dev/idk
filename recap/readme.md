# Recapitulation

## Quick Start

### Postgres

```go
package main

import (
	"bytes"
	"context"
	"github.com/ideatech-dev/idk/recap"
	"gorm.io/gorm"
	"log"
)

func main() {
	ctx := context.Background()
	writer := new(bytes.Buffer)

	file := recap.NewRecapExcelFile()
	defer file.Close()

	sheet1 := recap.NewSheet("sheet_1")
	sheet2 := recap.NewSheet("sheet_2")

	// insert source
	tableSchema := "public"
	tableName := "user"
	query := &gorm.DB{}
	
	err := sheet1.SourceGormPostgres(ctx, tableSchema, tableName, query)
	if err != nil {
		log.Fatal(err)
	}

	file.AddSheet(sheet1)
	file.AddSheet(sheet2)
	
	file.Write(writer)
}
```