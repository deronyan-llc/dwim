package ingest

import (
	"strings"

	"github.com/knakk/rdf"
)

type Class struct {
	Name       string
	Properties []Property
}

type Property struct {
	Name string
	Type string
}

func parseRDFSchema(schema string) ([]Class, error) {
	var classes []Class

	g := rdf.NewGraph()
	err := g.Parse(schema, rdf.Turtle)
	if err != nil {
		return classes, err
	}

	for _, class := range g.ResourcesOfType(rdf.OWL("Class")) {
		cls := Class{Name: class.Frag()}
		for _, prop := range g.SubjectsOf(rdf.RDF("type"), rdf.OWL("ObjectProperty")) {
			if r, ok := g.One(prop, rdf.RDFS("domain")).(rdf.NamedNode); ok && r == class {
				cls.Properties = append(cls.Properties, Property{Name: prop.Frag()})
			}
		}
		classes = append(classes, cls)
	}

	return classes, nil
}

func goStruct(class Class) string {
	var structBuilder strings.Builder
	structBuilder.WriteString("type ")
	structBuilder.WriteString(class.Name)
	structBuilder.WriteString(" struct {\n")

	for _, prop := range class.Properties {
		structBuilder.WriteString("\t")
		structBuilder.WriteString(prop.Name)
		structBuilder.WriteString(" ")
		structBuilder.WriteString(prop.Type)
		structBuilder.WriteString(" `json:\"")
		structBuilder.WriteString(strings.ToLower(prop.Name))
		structBuilder.WriteString("\" bson:\"")
		structBuilder.WriteString(strings.ToLower(prop.Name))
		structBuilder.WriteString("\"`\n")
	}

	structBuilder.WriteString("}\n\n")
	return structBuilder.String()
}

func generateGoSource(classes []Class) string {
	var sourceBuilder strings.Builder

	sourceBuilder.WriteString("package generated\n\n")

	for _, class := range classes {
		sourceBuilder.WriteString(goStruct(class))
	}

	return sourceBuilder.String()
}

// Implement the RDF schema parsing logic here.
// This function should return a list of Class instances based on the RDF schema.
func parseRDFSchema(schema string) ([]Class, error) {
	return []Class{}, nil
}
