package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/knakk/rdf"
)

type SchemaClass struct {
	Name       string
	Properties []SchemaProperty
}

type SchemaProperty struct {
	Name   string
	Domain string
	Range  string
}

// parseRDFSchema reads the Turtle RDF schema and returns a map of SchemaClass.
func parseRDFSchema(file string) (map[string]SchemaClass, error) {
	// Read the Turtle RDF schema file.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)

	// Parse the Turtle RDF schema.
	triples, err := rdf.NewTripleDecoder(reader, rdf.Turtle).DecodeAll()
	if err != nil {
		return nil, err
	}

	// Convert the triples to an intermediate representation.
	classes := make(map[string]SchemaClass)
	for _, triple := range triples {
		// Check if the triple describes a class or a property.
		switch triple.Pred.String() {
		case "http://www.w3.org/1999/02/22-rdf-syntax-ns#type":
			if triple.Obj.String() == "http://www.w3.org/2000/01/rdf-schema#Class" {
				classes[triple.Subj.String()] = SchemaClass{Name: triple.Subj.String()}
			}
		case "http://www.w3.org/2000/01/rdf-schema#domain":
			if class, ok := classes[triple.Obj.String()]; ok {
				property := SchemaProperty{
					Name:   triple.Subj.String(),
					Domain: triple.Obj.String(),
				}
				class.Properties = append(class.Properties, property)
				classes[triple.Obj.String()] = class
			}
		case "http://www.w3.org/2000/01/rdf-schema#range":
			for _, class := range classes {
				for i, property := range class.Properties {
					if property.Name == triple.Subj.String() {
						class.Properties[i].Range = triple.Obj.String()
						classes[triple.Obj.String()] = class
						break
					}
				}
			}
		}
	}

	return classes, nil
}

// generateGoCode generates GoLang code based on the given SchemaClass map.
func generateGoCode(classes map[string]SchemaClass) {
	const tmpl = `package main

type {{ .Name }} struct {
{{- range .Properties }}
	{{ .Name }} {{ .Range }}
{{- end }}
}
`

	t := template.Must(template.New("class").Parse(tmpl))

	// Generate GoLang code for each class in the map.
	for _, class := range classes {
		// Generate a Go source file for the class.
		filename := fmt.Sprintf("%s.go", strings.ToLower(class.Name))
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer file.Close()

		// Execute the template with the class data.
		err = t.Execute(file, class)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
	}
}

func main() {
	// Parse the RDF schema file.
	classes, err := parseRDFSchema("../schemas/ethereum/ethereum-core.ttl")
	if err != nil {
		fmt.Println("Error parsing RDF schema:", err)
		os.Exit(1)
	}

	// Generate the GoLang code.
	generateGoCode(classes)
}
