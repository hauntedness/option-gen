package optiongen

// type CallOption func(*callOptions)
const templateOptionType = `type {{OptionTypeName .TypeName}} func(*{{.TypeName}})`

//	var WithIntField = func(intField int) CallOption {
//		return func(op *callOptions) {
//			op.intField = intField
//		}
//	}
const templateVariable = `
var {{ OptionVarName .FieldNameByIndex .WithPrefix .WithPostfix}} = func({{.ParamNameByIndex}} {{.ParamTypeByIndex}}) {{OptionTypeName .TypeName}} {
	return func(op *{{.TypeName}}) {
		op.{{.FieldNameByIndex}} = {{.ParamNameByIndex}}
	}
}`

// templateApplyFunc render a function that could apply options
//
//	func (op *callOptions) ApplyOptions(opts ...CallOption) {
//		for i := range opts {
//			opts[i](op)
//		}
//	}
const templateApplyFunc = `
func (op *{{.TypeName}}) ApplyOptions(opts ...{{OptionTypeName .TypeName}}) {
	for i := range opts {
		opts[i](op)
	}
}`
