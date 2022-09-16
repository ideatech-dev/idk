package recap

import (
	"github.com/xuri/excelize/v2"
	"os"
	"testing"
)

func TestFileExcel_Write(t *testing.T) {
	mockSheet := &Sheet{
		name:    "mock_sheet",
		columns: []string{"no", "name", "age"},
		data: []map[string]interface{}{
			{"no": 1, "name": "foo", "age": 10},
			{"no": 2, "name": "bar", "age": 15},
			{"no": 3, "name": "foobar", "age": 20},
		},
	}

	mockSheet2 := &Sheet{
		name:    "mock_sheet_2",
		columns: []string{"no", "name", "email"},
		data: []map[string]interface{}{
			{"no": 1, "name": "foo", "email": "foo@gmail.com"},
			{"no": 2, "name": "bar", "email": "bar@gmail.com"},
			{"no": 3, "name": "foobar", "email": "foobar@gmail.com"},
		},
	}

	type fields struct {
		file   *excelize.File
		sheets []*Sheet
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "generate single sheet excel",
			fields: fields{
				file:   excelize.NewFile(),
				sheets: []*Sheet{mockSheet},
			},
			wantErr: false,
		},
		{
			name: "generate double sheet excel",
			fields: fields{
				file:   excelize.NewFile(),
				sheets: []*Sheet{mockSheet, mockSheet2},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileExcel{
				file:   tt.fields.file,
				sheets: tt.fields.sheets,
			}

			writerFile, err := os.Create("../temp/TestFileExcel_Write_" + tt.name + ".xlsx")
			if err != nil {
				t.Error(err)
				return
			}
			defer writerFile.Close()

			err = f.Write(writerFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
