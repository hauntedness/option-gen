package optiongen

import (
	"log"
	"math"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
)

type option struct {
	postfix string
}

type Option func(o *option)

var WithPostfix = func(postfix string) func(o *option) {
	return func(o *option) {
		o.postfix = postfix
	}
}

func ExecuteAny(value any, args ...Option) string {
	return ExecuteType(reflect.TypeOf(value), args...)
}

func ExecuteType(typ reflect.Type, args ...Option) string {
	return ExecuteString(typ.Name(), typ.PkgPath(), args...)
}

func ExecuteString(typeName string, packagePath string, args ...Option) string {
	// load package
	g, err := LoadDefinition(packagePath, typeName, &packages.Config{Mode: math.MaxInt})
	if err != nil {
		log.Fatal(err)
	}
	var option option
	for i := range args {
		args[i](&option)
	}
	g.WithPostfix = option.postfix
	b := &strings.Builder{}
	// gen declare option type
	str := g.RenderOptionType()
	b.WriteString("\n")
	b.WriteString(str)
	// gen apply func
	b.WriteString("\n")
	str = g.RenderApplyFunc()
	b.WriteString(str)
	// gen options
	for i := range g.Fields {
		clone := g
		clone.Index = i
		b.WriteString("\n")
		str = clone.RenderOptionVariable()
		b.WriteString(str)
	}
	return b.String()
}
