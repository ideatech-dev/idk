package recap

import (
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"io"
	"strconv"
)

type IFileExcel interface {
	AddSheet(sheet *Sheet)
	Write(writer io.Writer) error
	Close()
}

var _ IFileExcel = &FileExcel{}

func NewRecapExcelFile() *FileExcel {
	return &FileExcel{
		file: excelize.NewFile(),
	}
}

type FileExcel struct {
	file   *excelize.File
	sheets []*Sheet
}

func (f *FileExcel) AddSheet(sheet *Sheet) {
	f.sheets = append(f.sheets, sheet)
}

func (f *FileExcel) Write(writer io.Writer) error {
	for _, sheet := range f.sheets {
		sheetName := sheet.name
		columns := sheet.columns

		f.file.NewSheet(sheetName)

		// process columns row
		if err := f.file.SetSheetRow(sheetName, "A1", &columns); err != nil {
			return errors.Wrapf(err, "error set Sheet columns row for sheet %s", sheetName)
		}

		// process data rows
		dataInterfaces := convertDataMapToArray2DInterface(sheet.data, columns)
		for i, dataRow := range dataInterfaces {
			axis := "A" + strconv.Itoa(i+2) // Start from A2
			if err := f.file.SetSheetRow(sheetName, axis, &dataRow); err != nil {
				return errors.Wrapf(err, "error set Sheet data row axis %s", axis)
			}
		}
	}

	// delete default Sheet (unused)
	if len(f.sheets) > 0 {
		f.file.DeleteSheet("Sheet1")
	}

	if err := f.file.Write(writer); err != nil {
		return errors.Wrapf(err, "error while write excel")
	}

	return nil
}

func (f *FileExcel) Close() {
	f.file.Close()
}
