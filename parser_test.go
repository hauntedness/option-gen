package optiongen

import (
	"math"
	"reflect"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestLoadDefinition(t *testing.T) {
	type args struct {
		packagePath string
		typeName    string
	}
	tests := []struct {
		name    string
		args    args
		wantG   Gen
		wantErr bool
	}{
		{
			name: "github.com/hauntedness/optiongen/internal",
			args: args{
				packagePath: "github.com/hauntedness/optiongen/internal",
				typeName:    "callOptions",
			},
			wantG: Gen{
				TypeName: "callOptions",
				Fields: []Field{
					{
						FieldName: "intField",
						FieldType: "int",
					},
					{
						FieldName: "stringField",
						FieldType: "string",
					},
					{
						FieldName: "interfaceField",
						FieldType: "interface{}",
					},
					{
						FieldName: "writer",
						FieldType: "io.Writer",
					},
				},
				Index:       0,
				WithPostfix: "",
				PackageName: "internal",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotG, err := LoadDefinition(tt.args.packagePath, tt.args.typeName, &packages.Config{Mode: math.MaxInt})
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadDefinition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotG, tt.wantG) {
				t.Errorf("LoadDefinition() = %v, want %v", gotG, tt.wantG)
			}
		})
	}
}
