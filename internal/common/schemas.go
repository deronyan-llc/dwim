package common

import "github.com/knakk/rdf"

type SchemaClass struct {
	Name       string
	Properties []SchemaProperty
	TargetDir  string
	Package    string
	Imports    []string
}

type SchemaProperty struct {
	Name     string
	Domain   string
	Range    string
	LangType string
}

type SchemaContext struct {
	Path         string     // file path or uri pointing to the schema
	StrippedPath string     // for file names, strips off the file extension
	Format       rdf.Format // the RDF format of the schema
	Extension    string     // the file extension of the schema
	Classes      map[string]*SchemaClass
}
