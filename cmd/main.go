package main

import (
	"flag"
	"fmt"

	"github.com/hauntedness/optiongen"
)

func main() {
	typeName := flag.String("o", "", "the type name to gen options")
	packagePath := flag.String("p", ".", "the package name to gen options, default to current package")
	varPostfix := flag.String("t", "", `specify the variables postfix, default ""`)
	flag.Parse()
	generated := optiongen.ExecuteString(*typeName, *packagePath, optiongen.WithPostfix(*varPostfix))
	fmt.Println(generated)
}
