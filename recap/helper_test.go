package recap

import (
	"reflect"
	"testing"
)

func Test_convertDataMapToArray2DInterface(t *testing.T) {
	type args struct {
		data []map[string]interface{}
		cols []string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput [][]interface{}
	}{
		{
			name: "positive - all columns same",
			args: args{
				data: []map[string]interface{}{
					{"name": "John", "age": 20},
					{"name": "Doe", "age": 21},
				},
				cols: []string{"name", "age"},
			},
			wantOutput: [][]interface{}{
				{"John", 20},
				{"Doe", 21},
			},
		},
		{
			name: "positive - partial columns",
			args: args{
				data: []map[string]interface{}{
					{"name": "John", "age": 20},
					{"name": "Doe", "age": 21},
				},
				cols: []string{"name"},
			},
			wantOutput: [][]interface{}{
				{"John"},
				{"Doe"},
			},
		},
		{
			name: "positive - partial column not same",
			args: args{
				data: []map[string]interface{}{
					{"name": "John", "age": 20},
					{"name": "Doe", "age": 21},
				},
				cols: []string{"name", "unknown"},
			},
			wantOutput: [][]interface{}{
				{"John", nil},
				{"Doe", nil},
			},
		},
		{
			name: "positive - all column not same",
			args: args{
				data: []map[string]interface{}{
					{"name": "John", "age": 20},
					{"name": "Doe", "age": 21},
				},
				cols: []string{"unknown1"},
			},
			wantOutput: [][]interface{}{
				{nil},
				{nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := convertDataMapToArray2DInterface(tt.args.data, tt.args.cols); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("convertDataMapToArray2DInterface() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

