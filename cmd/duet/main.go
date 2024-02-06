package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Class represents an RDF class.
type Class struct {
	Name       string
	Properties []*Property
}

// Property represents an RDF property.
type Property struct {
	Name        string
	Type        string
	Description string
}

// File represents a generated Go file.
type File struct {
	Name        string
	PackageName string
	Structs     []*Struct
	Imports     []string
}

// Struct represents a generated Go struct.
type Struct struct {
	Name   string
	Fields []*Field
}

// Field represents a generated Go field.
type Field struct {
	Name        string
	Type        string
	Description string
}

func main() {
	// Get the input file names from the command line.
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run generate.go <input_file1> <input_file2> ...")
	}
	inputFiles := os.Args[1:]

	// Read the input files.
	var inputData [][]byte
	for _, inputFile := range inputFiles {
		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		inputData = append(inputData, data)
	}

	// Parse the input files.
	var rdf []*interface{}
	for _, data := range inputData {
		var r interface{}
		err := json.Unmarshal(data, &r)
		if err != nil {
			log.Fatal(err)
		}
		rdf = append(rdf, &r)
	}

	// Generate the Go files.
	var classes []*Class
	for _, r := range rdf {
		for _, c := range (*r).([]interface{}) {
			class := &Class{
				Name:       c.(map[string]interface{})["@type"].(string),
				Properties: []*Property{},
			}
			for _, p := range c.(map[string]interface{})["properties"].(map[string]interface{}) {
				property := &Property{
					Name:        p.(map[string]interface{})["name"].(string),
					Type:        p.(map[string]interface{})["type"].(string),
					Description: p.(map[string]interface{})["description"].(string),
				}
				class.Properties = append(class.Properties, property)
			}
			classes = append(classes, class)
		}
	}

	var files []*File
	for _, c := range classes {
		file := &File{
			Name:        c.Name + ".go",
			PackageName: "foaf",
			Structs:     []*Struct{},
			Imports:     []string{},
		}
		for _, p := range c.Properties {
			structField := &Field{
				Name:        p.Name,
				Type:        p.Type,
				Description: p.Description,
			}
			file.Structs[0].Fields = append(file.Structs[0].Fields, structField)
		}
		files = append(files, file)
	}

	// Write the Go files.
	for _, f := range files {
		err := ioutil.WriteFile(f.Name, []byte(generateGoFile(f)), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func generateGoFile(f *File) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("package %s\n\n", f.PackageName))

	for _, s := range f.Structs {
		sb.WriteString(fmt.Sprintf("type %s struct {\n", s.Name))
		for _, f := range s.Fields {
			sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"` // %s\n", f.Name, f.Type, f.Name, f.Description))
		}
		sb.WriteString("}\n\n")
	}

	for _, i := range f.Imports {
		sb.WriteString(fmt.Sprintf("import \"%s\"\n", i))
	}

	return sb.String()
}
