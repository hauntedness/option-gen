package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hauntedness/optiongen"
)

func main() {
	typeName := flag.String("o", "", "the type name to gen options")
	packagePath := flag.String("p", ".", "the package name to gen options, default to current package")
	varPostfix := flag.String("t", "", `specify the variables postfix, default ""`)
	flag.Parse()
	if *typeName == "" {
		log.Fatal("type name is mandatory: optiongen -o someoption")
	}
	generated := optiongen.ExecuteString(*typeName, *packagePath, optiongen.WithPostfix(*varPostfix))
	fmt.Println(generated)
}
