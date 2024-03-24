package optiongen

import (
	"strings"
	"text/template"
)

// Gen is a wide type which will be rendered with a bunch of templates
type Gen struct {
	TypeName    string
	Fields      []Field
	Index       int
	WithPostfix string
	PackageName string
}

type Field struct {
	FieldName string
	FieldType string
}

func (g Gen) FieldNameByIndex() string {
	return g.Fields[g.Index].FieldName
}

func (g Gen) ParamNameByIndex() string {
	return strings.ToLower(g.Fields[g.Index].FieldName[:1]) + g.Fields[g.Index].FieldName[1:]
}

func (g Gen) FieldTypeByIndex() string {
	return g.Fields[g.Index].FieldType
}

func (g Gen) ParamTypeByIndex() string {
	//
	typ := g.Fields[g.Index].FieldType
	typ = strings.Replace(typ, "[]", "...", 1)
	return typ
}

func (g Gen) RenderOptionType() string {
	return g.Render(templateOptionType)
}

func (g Gen) RenderApplyFunc() string {
	return g.Render(templateApplyFunc[1:])
}

func (g Gen) RenderOptionVariable() string {
	return g.Render(templateVariable[1:])
}

func (g Gen) Render(tmpl string) string {
	t := template.New("f").Funcs(template.FuncMap{
		"OptionTypeName": OptionTypeName,
		"OptionVarName":  OptionVarName,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		panic(err)
	}
	b := &strings.Builder{}
	err = t.Execute(b, g)
	if err != nil {
		panic(err)
	}
	return b.String()
}
