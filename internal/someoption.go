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

var DefaultCallOption = &callOptions{}

type CallOption func(*callOptions)

var WithIntField = func(intField int) CallOption {
	return func(c *callOptions) {
		c.intField = intField
	}
}

var WithStringField = func(stringField string) CallOption {
	return func(c *callOptions) {
		c.stringField = stringField
	}
}

var WithInterfaceField = func(interfaceField interface{}) CallOption {
	return func(c *callOptions) {
		c.interfaceField = interfaceField
	}
}

var WithWriter = func(writer io.Writer) CallOption {
	return func(c *callOptions) {
		c.writer = writer
	}
}

var WithNumber = func(number json.Number) CallOption {
	return func(c *callOptions) {
		c.number = number
	}
}

func (c *callOptions) Apply(opts ...CallOption) {
	for i := range opts {
		opts[i](c)
	}
}

func (c *callOptions) IntField(intField int) *callOptions {
	c.intField = intField
	return c
}

func (c *callOptions) StringField(stringField string) *callOptions {
	c.stringField = stringField
	return c
}

func (c *callOptions) InterfaceField(interfaceField interface{}) *callOptions {
	c.interfaceField = interfaceField
	return c
}

func (c *callOptions) Writer(writer io.Writer) *callOptions {
	c.writer = writer
	return c
}

func (c *callOptions) Number(number json.Number) *callOptions {
	c.number = number
	return c
}
