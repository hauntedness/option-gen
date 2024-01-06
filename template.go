package main

// type CallOption func(*callOptions)
const templateOptionType = `type {{OptionTypeName .TypeName}} func(*{{.TypeName}})`

//	var WithIntField = func(intField int) CallOption {
//		return func(op *callOptions) {
//			op.intField = intField
//		}
//	}
const templateVariable = `
var {{ OptionVarName .FieldNameByIndex .WithPostfix}} = func({{.FieldNameByIndex}} {{.FieldTypeByIndex}}) {{OptionTypeName .TypeName}} {
	return func(op *{{.TypeName}}) {
		op.{{.FieldNameByIndex}} = {{.FieldNameByIndex}}
	}
}`

// templateApplyFunc render a function that could apply options
//
//	func (op *callOptions) ApplyOption(opts ...CallOption) {
//		for i := range opts {
//			opts[i](op)
//		}
//	}
const templateApplyFunc = `
func (op *{{.TypeName}}) ApplyOption(opts ...{{OptionTypeName .TypeName}}) {
	for i := range opts {
		opts[i](op)
	}
}`
