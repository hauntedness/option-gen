package optiongen

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

type option struct {
	prefix      string
	postfix     string
	writeFile   string
	autoImports bool
	builderMode bool
}

type Option func(o *option)

var WithPrefix = func(prefix string) func(o *option) {
	return func(o *option) {
		o.prefix = prefix
	}
}

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

var WithAutoImports = func(autoImports bool) func(o *option) {
	return func(o *option) {
		o.autoImports = autoImports
	}
}

var WithBuilderMode = func(builderMode bool) func(o *option) {
	return func(o *option) {
		o.builderMode = builderMode
	}
}

func ExecuteAny(value any, args ...Option) string {
	return ExecuteType(reflect.TypeOf(value), args...)
}

func ExecuteType(typ reflect.Type, args ...Option) string {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
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
	g.WithPrefix = option.prefix
	b := &bytes.Buffer{}

	b.WriteString("package ")
	b.WriteString(g.PackageName)
	b.WriteString("\n\n\n")
	// gen declare option type
	if !option.builderMode {
		str := g.RenderOptionType()
		b.WriteString("\n\n")
		b.WriteString(str)
	}
	// gen apply func
	b.WriteString("\n\n")
	if !option.builderMode {
		str := g.RenderApplyFunc()
		b.WriteString(str)
	}
	// gen options
	for i := range g.Fields {
		clone := g
		clone.Index = i
		b.WriteString("\n\n")
		if option.builderMode {
			str := clone.RenderChainFunc()
			b.WriteString(str)
		} else {
			str := clone.RenderOptionVariable()
			b.WriteString(str)
		}
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

func ExecuteAll(values []any, args ...Option) string {
	var types []reflect.Type
	for _, value := range values {
		types = append(types, reflect.TypeOf(value))
	}
	return ExecuteAllType(types, args...)
}

func ExecuteAllType(types []reflect.Type, args ...Option) string {
	var typeNames []string
	var pkg string
	for _, typ := range types {
		for typ.Kind() == reflect.Pointer {
			typ = typ.Elem()
		}
		typeNames = append(typeNames, typ.Name())
		if pkg != "" && typ.PkgPath() != pkg {
			log.Panic("types are not in same package")
		}
		pkg = typ.PkgPath()
	}
	return ExecuteAllString(typeNames, pkg, args...)
}

func ExecuteAllString(typeNames []string, packagePath string, args ...Option) string {
	// load package
	gs, err := LoadDefinitions(packagePath, typeNames, &packages.Config{Mode: math.MaxInt})
	if err != nil {
		log.Panic(err)
	}
	if len(gs) == 0 {
		log.Panic(fmt.Errorf("require at least 1 type name."))
	}
	b := &bytes.Buffer{}
	option := option{autoImports: true, builderMode: true}
	for i := range args {
		args[i](&option)
	}
	if !option.builderMode {
		log.Panic(fmt.Errorf("Only support builder mode for batch generation."))
	}
	b.WriteString("package " + gs[0].PackageName)
	b.WriteString("\n\n\n")
	for _, g := range gs {
		g.WithPostfix = option.postfix
		g.WithPrefix = option.prefix

		// gen apply func
		b.WriteString("\n\n")

		// gen options
		for i := range g.Fields {
			clone := g
			clone.Index = i
			b.WriteString("\n\n")
			str := clone.RenderChainFunc()
			b.WriteString(str)
		}
		b.WriteString("\n\n\n")
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
