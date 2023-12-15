package main

import (
	"fmt"
	"log"
	"os"

	"deronyan.com/columbo/internal/parsers"
	"deronyan.com/columbo/internal/src_generators"
)

func main() {
	// Parse the RDF schema files
	// read all schema files in the schemas directory
	dir := os.Args[1]
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var parser parsers.Parser = parsers.RDFParser{}
	var goSrcGenerator src_generators.SrcGenerator = src_generators.GoSrcGenerator{}

	for _, file := range files {
		fmt.Printf("Parsing RDF schema for file(%s)\n", file.Name())
		context, err := parser.Parse(dir + file.Name())
		if err != nil {
			fmt.Printf("Error parsing RDF schema for file(%s): %v\n", file.Name(), err)
			continue
		}

		// Generate the GoLang code.
		if err := goSrcGenerator.Generate(context); err != nil {
			fmt.Printf("Error generating GoLang code for file(%s): %v\n", file.Name(), err)
			continue
		}
	}
}
