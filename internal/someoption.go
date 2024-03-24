package internal

import (
	"encoding/json"
	"io"
)

type SomeType struct{}

type callOptions struct {
	intField       int
	stringField    string
	interfaceField interface{}
	writer         io.Writer
	number         json.Number
}

var DefaultCallOption = callOptions{}

type CallOption func(*callOptions)

var WithIntField = func(intField int) CallOption {
	return func(op *callOptions) {
		op.intField = intField
	}
}

var WithStringField = func(stringField string) CallOption {
	return func(op *callOptions) {
		op.stringField = stringField
	}
}

var WithInterfaceField = func(interfaceField interface{}) CallOption {
	return func(op *callOptions) {
		op.interfaceField = interfaceField
	}
}

var WithWriter = func(writer io.Writer) CallOption {
	return func(op *callOptions) {
		op.writer = writer
	}
}

var WithNumber = func(number json.Number) CallOption {
	return func(op *callOptions) {
		op.number = number
	}
}

func (op *callOptions) Apply(opts ...CallOption) {
	for i := range opts {
		opts[i](op)
	}
}
