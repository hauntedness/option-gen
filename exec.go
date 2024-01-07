package optiongen

import (
	"log"
	"math"
	"os"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
)

type option struct {
	postfix   string
	writeFile string
}

type Option func(o *option)

var WithPostfix = func(postfix string) func(o *option) {
	return func(o *option) {
		o.postfix = postfix
	}
}

var WithWriteFile = func(writeFile string) func(o *option) {
	return func(o *option) {
		o.writeFile = writeFile
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
		log.Panic(err)
	}
	var option option
	for i := range args {
		args[i](&option)
	}
	g.WithPostfix = option.postfix
	b := &strings.Builder{}
	// gen declare option type
	str := g.RenderOptionType()
	b.WriteString("\n\n")
	b.WriteString(str)
	// gen apply func
	b.WriteString("\n\n")
	str = g.RenderApplyFunc()
	b.WriteString(str)
	// gen options
	for i := range g.Fields {
		clone := g
		clone.Index = i
		b.WriteString("\n\n")
		str = clone.RenderOptionVariable()
		b.WriteString(str)
	}
	content := b.String()
	if option.writeFile != "" {
		err := os.WriteFile(option.writeFile, []byte(content), os.ModePerm)
		if err != nil {
			log.Panic(err)
		}
	}
	return content
}
