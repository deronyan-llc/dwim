package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/deiu/rdf2go"
)

func main() {
	// Read RDF model file
	modelPath := os.Args[1]
	model, err := rdf2go.ReadFile(modelPath)
	if err != nil {
		fmt.Println("Error reading RDF model:", err)
		os.Exit(1)
	}

	// Identify classes
	classes := map[string]*rdf2go.Resource{}
	for _, triple := range model.Triples() {
		subject := triple.Subject.Value
		predicate := triple.Predicate.Value
		obj := triple.Object.Value

		if predicate == rdf2go.RDF_TYPE {
			if strings.HasPrefix(obj, rdf2go.RDFS_CLASS_URI) {
				className := strings.TrimPrefix(obj, rdf2go.RDFS_CLASS_URI)
				classes[className] = model.Resource(subject)
			}
		}
	}

	// Generate GoLang files for each class
	for className, resource := range classes {
		outputFile := filepath.Join("models", className+".go")
		err := generateGolangFile(outputFile, className, resource, model)
		if err != nil {
			fmt.Println("Error generating GoLang file:", err)
			os.Exit(1)
		}
	}
}

func generateGolangFile(outputFile, className string, resource *rdf2go.Resource, model *rdf2go.Graph) error {
	content := fmt.Sprintf("package models\n\ntype %s struct {\n", className)

	// Generate fields for each property
	for _, triple := range model.TriplesMatching(resource, rdf2go.Predicate(nil), rdf2go.Object(nil)) {
		property := triple.Predicate.Value
		fieldType := getGolangType(model, property)
		content += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", strings.Title(property), fieldType, property)
	}

	content += "}\n"

	// Write content to file
	err := ioutil.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Generated GoLang file:", outputFile)
	return nil
}

func getGolangType(model *rdf2go.Graph, property string) string {
	// Implement logic to map RDF data types to GoLang types
	// This is a simplified example and needs further development
	switch property {
	case "http://www.w3.org/2001/XMLSchema#string":
		return "string"
	case "http://www.w3.org/2001/XMLSchema#int":
		return "int"
	default:
		return "interface{}"
	}
}
