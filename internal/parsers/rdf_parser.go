package parsers

import (
	"bytes"
	"os"
	"slices"

	"deronyan-llc.com/rdf"
	"deronyan.com/columbo/internal/common"
)

type RDFParser struct {
}

func (p RDFParser) Parse(file string) (*common.SchemaContext, error) {
	schemaPath, err := common.ParseSchemaPath(file)
	if err != nil {
		return nil, err
	}
	schemaContext := &common.SchemaContext{
		SchemaPath: *schemaPath,
		Classes:    make(map[string]*common.SchemaClass),
	}

	// Read the RDF schema file.
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)

	// Parse the RDF schema.
	decoder := rdf.NewTripleDecoder(reader, schemaContext.SchemaPath.RdfFormat)
	triples, err := decoder.DecodeAll()
	if err != nil {
		return nil, err
	}

	/*
	 * TODO - make this better
	 */
	for _, triple := range triples {
		// Check if the triple describes a class or a property.
		switch triple.Pred.String() {
		case "http://www.w3.org/1999/02/22-rdf-syntax-ns#type":
			if triple.Obj.String() == "http://www.w3.org/2000/01/rdf-schema#Class" {
				schemaContext.Classes[triple.Subj.String()] = &common.SchemaClass{Name: triple.Subj.String()}
			}
		case "http://www.w3.org/2000/01/rdf-schema#domain":
			if class, ok := schemaContext.Classes[triple.Obj.String()]; ok {
				property := &common.SchemaProperty{
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

						/*
						 * TODO - make this better, more abstract using a map of RDF types to Go types.
						 * Likely, some interface parent is needed for all the types.
						 */

						// assign the lang-specific data type for the Range
						objStr := triple.Obj.String()
						switch objStr {
						case common.XsdInteger.Str, common.XsdInt.Str: // "http://www.w3.org/2001/XMLSchema#integer":
							class.Properties[i].LangType = "int"
						case common.XsdString.Str: // "http://www.w3.org/2001/XMLSchema#string":
							class.Properties[i].LangType = "string"
						case common.XsdDateTime.Str: // "http://www.w3.org/2001/XMLSchema#dateTime":
							class.Properties[i].LangType = "time.Time"
							if !slices.Contains(class.Imports, "time") {
								class.Imports = append(class.Imports, "time")
							}
						default:
							// figure out how to resolve unknown types from other schemas
							// TODO: `Agent`, `anyURI`
							class.Properties[i].LangType = common.LocalName(triple.Obj.String())
							//class.Properties[i] = nil
						}
					}
				}
			}
		}
	}

	return schemaContext, nil
}
