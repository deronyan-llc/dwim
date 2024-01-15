package common

import (
	"fmt"
	"regexp"
	"strings"
)

func LocalName(uri string) string {
	parts := strings.Split(uri, "#")
	if len(parts) == 1 {
		parts = strings.Split(uri, "/")
	}
	return parts[len(parts)-1]
}

func SanitizeName(name string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	rtn := reg.ReplaceAllString(name, "")
	if rtn[0] >= '0' && rtn[0] <= '9' {
		rtn = "_" + rtn
	}
	return rtn
}

func FormatImports(imports []string) string {
	var result string
	for _, imp := range imports {
		result += fmt.Sprintf("\t\"%s\"\n", imp)
	}
	return result
}
