package recap

func convertDataMapToArray2DInterface(data []map[string]interface{}, cols []string) (output [][]interface{}) {
	for _, dataRow := range data {
		var outputRow []interface{}
		for _, col := range cols {
			outputRow = append(outputRow, dataRow[col])
		}
		output = append(output, outputRow)
	}

	return
}
