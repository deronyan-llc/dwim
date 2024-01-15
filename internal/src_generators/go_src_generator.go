package src_generators

import (
	"fmt"
	"go/format"
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/deronyan-llc/columbo/internal/common"
)

type GoSrcGenerator struct {
}

// combine all the imports and types from all the classes together
type GoStructs struct {
	Imports []string
	Classes []*common.SchemaClass
}

// generateGoCode generates GoLang code based on the given SchemaClass map.
func (g GoSrcGenerator) Generate(schemaContext *common.SchemaContext) error {
	const tmpl = `package generated
	
{{- if .Imports }}

import (
{{ .Imports | imports }}
)


{{- end }}

{{- if .Classes }}
	{{- range .Classes }}

// {{ .Name | localName | sanitizeName | title }} is a generated struct representing the {{ .Name }} class.
type {{ .Name | localName | sanitizeName | title }} struct {
		{{- if .Properties }}
			{{- range .Properties }}
	{{ .Name | localName | sanitizeName | title }} {{ .LangType }} ` +
		"`json:\"{{ .Name | localName | sanitizeName | lower }}\"`" +
		"{{ .Comment }}" + `
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
	goSrcFileName := fmt.Sprintf("%s/%s.go",
		schemaContext.OutputPath,
		schemaContext.SchemaPath.StrippedFileName,
	)

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

	// execute the template
	stringBuf := strings.Builder{}
	err = t.Execute(&stringBuf, goStructs)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return err
	}

	// create the output directory if it doesn't exist
	os.MkdirAll(schemaContext.OutputPath, os.ModePerm)
	file, err := os.Create(goSrcFileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// write the generated Go source to a file, unformatted initially
	_, err = file.Write([]byte(stringBuf.String()))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	// format the Go source code
	formatted, err := format.Source([]byte(stringBuf.String()))
	if err != nil {
		fmt.Println("Error formatting source:", err)
		return err
	}

	// write the formatted Go source to a file
	_, err = file.WriteAt(formatted, 0)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}
