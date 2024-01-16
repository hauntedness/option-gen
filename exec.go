package optiongen

import (
	"bytes"
	"log"
	"math"
	"os"
	"reflect"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

type option struct {
	postfix     string
	writeFile   string
	autoImports bool
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

var WithAutoimports = func(autoImports bool) func(o *option) {
	return func(o *option) {
		o.autoImports = autoImports
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
	option := option{autoImports: true}
	for i := range args {
		args[i](&option)
	}
	g.WithPostfix = option.postfix
	b := &bytes.Buffer{}

	b.WriteString("package ")
	b.WriteString(g.PackageName)
	b.WriteString("\n\n\n")
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
	content := b.Bytes()
	if option.autoImports {
		content, err = imports.Process("", content, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	if option.writeFile != "" {
		err := os.WriteFile(option.writeFile, content, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}
	}
	return string(content)
}
