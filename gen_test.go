package main

import (
	"strings"
	"testing"
)

func TestGen_RenderOptionType(t *testing.T) {
	type args struct {
		g Gen
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "callConfig",
			args: args{
				g: Gen{TypeName: "callConfig"},
			},
			want: "type CallOption func(*callConfig)",
		},
		{
			name: "call",
			args: args{
				g: Gen{TypeName: "call"},
			},
			want: "type CallOption func(*call)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.g.RenderOptionType(); got != tt.want {
				t.Errorf("Gen.RenderOptionType() = %v, want %v", got, tt.want)
			}
		})
	}
}

var t1 = `
func (op *callConfig) ApplyOption(opts ...CallOption) {
	for i := range opts {
		opts[i](op)
	}
}`

func TestGen_RenderApplyFunc(t *testing.T) {
	type args struct {
		g Gen
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "callConfig",
			args: args{
				g: Gen{TypeName: "callConfig"},
			},
			want: strings.TrimLeft(t1, "\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.g.RenderApplyFunc(); got != tt.want {
				t.Errorf("Gen.RenderApplyFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

const t2 = `
var WithSomeInt = func(someInt int) CallOption {
	return func(op *callConfig) {
		op.someInt = someInt
	}
}`

const t3 = `
var WithSomeIntSomehow = func(someInt int) CallOption {
	return func(op *callConfig) {
		op.someInt = someInt
	}
}`

func TestGen_RenderOptionVariable(t *testing.T) {
	type args struct {
		g Gen
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "callConfig",
			args: args{
				g: Gen{
					TypeName: "callConfig",
					Fields: []Field{
						{
							FieldName: "someField",
							FieldType: "string",
						},
						{
							FieldName: "someInt",
							FieldType: "int",
						},
					},
					Index: 1,
				},
			},
			want: strings.TrimLeft(t2, "\n"),
		},
		{
			name: "callConfig",
			args: args{
				g: Gen{
					TypeName: "callConfig",
					Fields: []Field{
						{
							FieldName: "someField",
							FieldType: "string",
						},
						{
							FieldName: "someInt",
							FieldType: "int",
						},
					},
					Index:       1,
					WithPostfix: "Somehow",
				},
			},
			want: strings.TrimLeft(t3, "\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.g.RenderOptionVariable(); got != tt.want {
				t.Errorf("Gen.RenderApplyFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
