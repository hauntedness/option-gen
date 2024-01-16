package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hauntedness/optiongen"
)

func main() {
	typeName := flag.String("type", "", "the type name to gen options")
	packagePath := flag.String("package", ".", "the package name to gen options, default to current package")
	postfix := flag.String("postfix", "", `specify the variables postfix, default ""`)
	writeFile := flag.String("writeFile", "", "specify which file writes to, default no write")
	autoImport := flag.Bool("autoImport", true, "organize import automatically")
	flag.Parse()
	if *typeName == "" {
		log.Fatal("type name is mandatory: optiongen -o someoption")
	}
	opts := []optiongen.Option{optiongen.WithPostfix(*postfix), optiongen.WithAutoimports(*autoImport), optiongen.WithWriteFile(*writeFile)}
	generated := optiongen.ExecuteString(*typeName, *packagePath, opts...)
	fmt.Println(generated)
}
