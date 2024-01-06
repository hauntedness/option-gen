package main

import (
	"testing"
)

func TestOptionTypeName(t *testing.T) {
	type args struct {
		applyeeName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "callOptions",
			args: args{
				applyeeName: "callOptions",
			},
			want: "CallOption",
		},
		{
			name: "callConfig",
			args: args{
				applyeeName: "callConfig",
			},
			want: "CallOption",
		},
		{
			name: "callConfigs",
			args: args{
				applyeeName: "callConfigs",
			},
			want: "CallOption",
		},
		{
			name: "call",
			args: args{
				applyeeName: "call",
			},
			want: "CallOption",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptionTypeName(tt.args.applyeeName); got != tt.want {
				t.Errorf("OptionTypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionVarName(t *testing.T) {
	type args struct {
		fieldName string
		postfix   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "intField no postfix",
			args: args{
				fieldName: "intField",
			},
			want: "WithIntField",
		},
		{
			name: "postfix",
			args: args{
				fieldName: "StringField",
				postfix:   "Option",
			},
			want: "WithStringFieldOption",
		},
		{
			name: "underline",
			args: args{
				fieldName: "_intField",
			},
			want: "With_intField",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptionVarName(tt.args.fieldName, tt.args.postfix); got != tt.want {
				t.Errorf("OptionVarName() = %v, want %v", got, tt.want)
			}
		})
	}
}
