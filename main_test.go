package main

import (
	"fmt"
	"testing"

	"github.com/hauntedness/option-gen/internal"
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
			name: "github.com/hauntedness/option-gen/internal",
			args: args{
				typeName:    "callOptions",
				packagePath: "github.com/hauntedness/option-gen/internal",
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
			name: "github.com/hauntedness/option-gen/internal",
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
