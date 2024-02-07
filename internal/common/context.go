package common

import (
	"fmt"
	"strings"

	"path/filepath"

	"github.com/deronyan-llc/rdf/rdf"
)

type SchemaClass struct {
	Term       rdf.Term
	TermString string
	Properties []*SchemaProperty
	Package    string
	Imports    []string
}

type SchemaProperty struct {
	Term       rdf.Term
	TermString string
	Domain     rdf.Term
	Range      rdf.Term
	GoLangType string
	Comment    string
}

type SchemaPath struct {
	Provided         string     // the original input path
	Path             string     // file path or uri pointing to the schema, without the file name+extension
	FileName         string     // for file names, includes extension
	StrippedFileName string     // for file names, without extension
	Extension        string     // the file extension of the schema
	RdfFormat        rdf.Format // the RDF format of the schema
}

type SchemaContext struct {
	SchemaPath  SchemaPath
	OutputPath  string
	PackageName string
	Classes     map[rdf.Term]*SchemaClass
}

var FileExtensionsToRdfFormats map[string]rdf.Format = map[string]rdf.Format{
	"ttl": rdf.Turtle,
	"xsd": rdf.RDFXML,
	"xml": rdf.RDFXML,
	"n3":  rdf.NTriples,
	"nq":  rdf.NQuads,
}

func ParseSchemaPath(schemaPathStr string) (*SchemaPath, error) {
	// for now, only assume `inputPath` is a file path, and extension is required to infer the RDF format
	// TODO: handle URL's as well

	dotPos := strings.LastIndex(schemaPathStr, ".")
	if dotPos == -1 {
		return nil, fmt.Errorf("invalid schema path(%s)", schemaPathStr)
	}

	schemaPath := &SchemaPath{
		Provided: schemaPathStr,
	}
	schemaPath.Path = filepath.Dir(schemaPathStr)
	schemaPath.Extension = filepath.Ext(schemaPathStr)
	schemaPath.FileName = filepath.Base(schemaPathStr)
	schemaPath.StrippedFileName = strings.TrimSuffix(schemaPath.FileName, schemaPath.Extension)

	if rdfFormat, ok := FileExtensionsToRdfFormats[schemaPath.Extension[1:]]; ok {
		schemaPath.RdfFormat = rdfFormat
	} else {
		return nil, fmt.Errorf("unsupported rdf format for schema path(%s)", schemaPath.Provided)
	}

	return schemaPath, nil
}
