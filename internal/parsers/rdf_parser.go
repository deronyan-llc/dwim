package parsers

import (
	"bytes"
	"fmt"
	"os"
	"slices"

	"deronyan.com/columbo/internal/common"
	"github.com/knakk/rdf"
)

type RDFParser struct {
}

func (p RDFParser) Parse(file string) (*common.SchemaContext, error) {
	// figure out what format of RDF file this is
	rdfFormat, _, _ := FileExtToRdfFormat(file)
	if rdfFormat == rdf.Format(-1) {
		return nil, fmt.Errorf("unknown RDF file format for file(%s)", file)
	}

	// Read the RDF schema file.
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)

	// Parse the RDF schema.
	decoder := rdf.NewTripleDecoder(reader, rdfFormat)
	triples, err := decoder.DecodeAll()
	if err != nil {
		return nil, err
	}

	rdfFormat, strippedFile, extension := FileExtToRdfFormat(file)
	schemaContext := &common.SchemaContext{
		Path:         file,
		StrippedPath: strippedFile,
		Format:       rdfFormat,
		Extension:    extension,
		Classes:      make(map[string]*common.SchemaClass),
	}

	for _, triple := range triples {
		// Check if the triple describes a class or a property.
		switch triple.Pred.String() {
		case "http://www.w3.org/1999/02/22-rdf-syntax-ns#type":
			if triple.Obj.String() == "http://www.w3.org/2000/01/rdf-schema#Class" {
				schemaContext.Classes[triple.Subj.String()] = &common.SchemaClass{Name: triple.Subj.String()}
			}
		case "http://www.w3.org/2000/01/rdf-schema#domain":
			if class, ok := schemaContext.Classes[triple.Obj.String()]; ok {
				property := common.SchemaProperty{
					Name:     triple.Subj.String(),
					Domain:   triple.Obj.String(),
					LangType: "NoRange",
				}
				class.Properties = append(class.Properties, property)
			}
		case "http://www.w3.org/2000/01/rdf-schema#range":
			for _, class := range schemaContext.Classes {
				for i, property := range class.Properties {
					if property.Name == triple.Subj.String() {
						class.Properties[i].Range = triple.Obj.String()
						class.Properties[i].LangType = triple.Obj.String()
						//classes[triple.Obj.String()] = class

						// assign the lang-specific data type for the Range
						switch triple.Obj.String() {
						case "http://www.w3.org/2001/XMLSchema#integer":
							class.Properties[i].LangType = "int"
						case "http://www.w3.org/2001/XMLSchema#string", "string", "xsd:string":
							class.Properties[i].LangType = "string"
						case "http://www.w3.org/2001/XMLSchema#dateTime":
							class.Properties[i].LangType = "time.Time"
							if !slices.Contains(class.Imports, "time") {
								class.Imports = append(class.Imports, "time")
							}
						default:
							class.Properties[i].LangType = common.LocalName(triple.Obj.String())
						}
					}
				}
			}
		}
	}

	return schemaContext, nil
}
