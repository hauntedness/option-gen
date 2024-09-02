package optiongen

import (
	"strings"
)

// OptionTypeName give the option type a proper name
// replace xxxOptions --> XXXOption
// replace xxxConfig --> XXXOption
// replace xxxConfigs --> XXXOption
func OptionTypeName(applyeeName string) string {
	newname := applyeeName
	if strings.HasSuffix(applyeeName, "Options") {
		newname = applyeeName[:len(applyeeName)-1]
	} else if strings.HasSuffix(applyeeName, "Config") {
		newname = applyeeName[:len(applyeeName)-len("Config")] + "Option"
	} else if strings.HasSuffix(applyeeName, "Configs") {
		newname = applyeeName[:len(applyeeName)-len("Configs")] + "Option"
	} else {
		newname = applyeeName + "Option"
	}
	return strings.ToTitle(newname[0:1]) + newname[1:]
}

// OptionVarName give the option variable a proper name
func OptionVarName(fieldName, prefix, postfix string) string {
	return "With" + prefix + strings.ToTitle(fieldName[0:1]) + fieldName[1:] + postfix
}

// BuilderFuncName give the input variable a proper name
func BuilderFuncName(fieldName, prefix, postfix string) string {
	return prefix + strings.ToTitle(fieldName[0:1]) + fieldName[1:] + postfix
}
