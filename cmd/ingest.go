package main

import (
	"io/ioutil"
	"os"
)

func main() {
	schemaFile := "path/to/your/schema.ttl"

	schema, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		fmt.Println("Error reading schema file:", err)
		return
	}

	classes, err := parseRDFSchema(string(schema))
	if err != nil {
		fmt.Println("Error parsing RDF schema:", err)
		return
	}

	goSource := generateGoSource(classes)
	fmt.Println(goSource)

	outputFile := "path/to/output/generated.go"
	err = ioutil.WriteFile(outputFile, []byte(goSource), 0644)
	if err != nil {
		fmt.Println("Error writing generated Go source file:", err)
		return
	}

	fmt.Println("Generated Go source file successfully:", outputFile)
}
