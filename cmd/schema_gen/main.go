package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deronyan-llc/columbo/internal/generators"
	"github.com/deronyan-llc/columbo/internal/generators/golang"
	"github.com/deronyan-llc/columbo/internal/parsers"
)

func main() {
	// Parse the RDF schema files
	// read all schema files in the schemas directory
	inputDir := os.Args[1]
	outputDir := os.Args[2]
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var parser parsers.Parser = parsers.RDFParser{}
	var goSrcGenerator generators.SrcGenerator = golang.GoSrcGenerator{}

	for _, file := range files {
		fmt.Printf("Parsing RDF schema for file(%s)\n", file.Name())
		context, err := parser.Parse(inputDir + "/" + file.Name())
		if err != nil {
			fmt.Printf("Error parsing RDF schema for file(%s): %v\n", file.Name(), err)
			continue
		}

		context.OutputPath = outputDir
		if err := goSrcGenerator.Generate(context); err != nil {
			fmt.Printf("Error generating GoLang code for file(%s): %v\n", file.Name(), err)
			continue
		}
	}
}
