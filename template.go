package optiongen

// type CallOption func(*callOptions)
const templateOptionType = `type {{OptionTypeName .TypeName}} func(*{{.TypeName}})`

//	var WithIntField = func(intField int) CallOption {
//		return func(c *callOptions) {
//			c.intField = intField
//		}
//	}
const templateVariable = `
var {{ OptionVarName .FieldNameByIndex .WithPrefix .WithPostfix}} = func({{.ParamNameByIndex}} {{.ParamTypeByIndex}}) {{OptionTypeName .TypeName}} {
	return func({{.ReceiverName}} *{{.TypeName}}) {
		{{.ReceiverName}}.{{.FieldNameByIndex}} = {{.ParamNameByIndex}}
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
func ({{.ReceiverName}} *{{.TypeName}}) ApplyOptions(opts ...{{OptionTypeName .TypeName}}) {
	for i := range opts {
		opts[i]({{.ReceiverName}})
	}
}`

// templateChainFunc render a function that could chain the property setters.
//
//	func (c *CallConfig) Name(name string) *CallConfig {
//		c.name = name
//		return c
//	}
const templateChainFunc = `
func ({{.ReceiverName}} *{{.TypeName}}) {{BuilderFuncName .FieldNameByIndex .WithPrefix .WithPostfix}}({{.ParamNameByIndex}} {{.ParamTypeByIndex}}) *{{.TypeName}} {
	{{.ReceiverName}}.{{.FieldNameByIndex}} = {{.ParamNameByIndex}}
	return {{.ReceiverName}}
}`
