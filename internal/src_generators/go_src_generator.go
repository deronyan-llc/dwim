package src_generators

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"text/template"

	"deronyan.com/columbo/internal/common"
)

type GoSrcGenerator struct {
}

// generateGoCode generates GoLang code based on the given SchemaClass map.
func (g GoSrcGenerator) Generate(schemaContext *common.SchemaContext) error {
	const tmpl = `package main
{{- if .Imports }}

import (
{{ .Imports | imports }}
)


{{- end }}

{{- if .Classes }}
	{{- range .Classes }}
type {{ .Name | localName | sanitizeName | title }} struct {
		{{- if .Properties }}
			{{- range .Properties }}
	{{ .Name | localName | sanitizeName | title }} {{ .LangType }} ` +
		"`json:\"{{ .Name | localName | sanitizeName | lower }}\"`" + `
			{{- end }}
		{{- end }}
}

	{{- end }}
{{- end }}
`
	funcMap := template.FuncMap{
		"title":        strings.Title,
		"lower":        strings.ToLower,
		"localName":    common.LocalName,
		"sanitizeName": common.SanitizeName,
		"imports":      common.FormatImports,
	}

	t, err := template.New("class").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	// Generate a Go source file for the classes.
	generatedGoSrcName := fmt.Sprintf("%s.go", strings.ToLower(schemaContext.StrippedPath))
	file, err := os.Create("../../gen/schemas/" + generatedGoSrcName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// combine all the imports and types from all the classes together
	type GoStructs struct {
		Imports []string
		Classes []*common.SchemaClass
	}
	goStructs := GoStructs{
		Imports: []string{},
		Classes: []*common.SchemaClass{},
	}
	for _, classes := range schemaContext.Classes {
		for _, importName := range classes.Imports {
			if !slices.Contains(goStructs.Imports, importName) {
				goStructs.Imports = append(goStructs.Imports, importName)
			}
		}
		goStructs.Classes = append(goStructs.Classes, classes)
	}

	err = t.Execute(file, goStructs)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return err
	}

	return nil
}
