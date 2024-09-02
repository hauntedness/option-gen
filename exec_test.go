package optiongen

import (
	"fmt"
	"testing"

	"github.com/hauntedness/optiongen/internal"
)

func TestExecuteString(t *testing.T) {
	type args struct {
		typeName    string
		packagePath string
		args        []Option
	}
	tests := []struct {
		name string
		args args
		want func(string) bool
	}{
		{
			name: "github.com/hauntedness/optiongen/internal",
			args: args{
				typeName:    "callOptions",
				packagePath: "github.com/hauntedness/optiongen/internal",
				args:        []Option{},
			},
			want: func(got string) bool {
				return got != ""
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteString(tt.args.typeName, tt.args.packagePath, tt.args.args...); !tt.want(got) {
				t.Errorf("ExecuteString() = %v", got)
			}
		})
	}
}

func TestExecuteAny(t *testing.T) {
	type args struct {
		value any
		args  []Option
	}
	tests := []struct {
		name string
		args args
		want func(string) bool
	}{
		{
			name: "github.com/hauntedness/optiongen/internal",
			args: args{
				value: internal.DefaultCallOption,
			},
			want: func(got string) bool {
				return got != ""
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteAny(tt.args.value, tt.args.args...); !tt.want(got) {
				t.Errorf("ExecuteAny() = %v", got)
			} else {
				fmt.Println(got)
			}
		})
	}
}

func TestExecuteAnyBuilder(t *testing.T) {
	type args struct {
		value any
		args  []Option
	}
	tests := []struct {
		name string
		args args
		want func(string) bool
	}{
		{
			name: "github.com/hauntedness/optiongen/internal",
			args: args{
				value: internal.DefaultCallOption,
				args:  []Option{WithBuilderMode(true)},
			},
			want: func(got string) bool {
				return got != ""
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteAny(tt.args.value, tt.args.args...); !tt.want(got) {
				t.Errorf("ExecuteAny() = %v", got)
			} else {
				fmt.Println(got)
			}
		})
	}
}
