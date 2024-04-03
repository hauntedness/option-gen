package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hauntedness/optiongen"
)

func main() {
	typeName := flag.String("type", "", "the type name to generate options on, mandatory")
	packagePath := flag.String("package", ".", "the package path to search types")
	prefix := flag.String("prefix", "", `specify the variables prefix, default ""`)
	postfix := flag.String("postfix", "", `specify the variables postfix, default ""`)
	writeFile := flag.String("writeFile", "", `specify which file writes to, default "no file"`)
	autoImports := flag.Bool("autoImports", true, "organize imports automatically")
	flag.Parse()
	if *typeName == "" {
		log.Fatal("type name is mandatory: optiongen --type someoption")
	}
	opts := []optiongen.Option{optiongen.WithPostfix(*postfix), optiongen.WithPrefix(*prefix), optiongen.WithAutoImports(*autoImports), optiongen.WithWriteFile(*writeFile)}
	generated := optiongen.ExecuteString(*typeName, *packagePath, opts...)
	fmt.Println(generated)
}
