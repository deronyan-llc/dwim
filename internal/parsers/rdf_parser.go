package parsers

import (
	"bytes"
	"os"
	"slices"

	"github.com/deronyan-llc/columbo/internal/common"
	"github.com/deronyan-llc/rdf/rdf"
)

type RDFParser struct {
}

func (p RDFParser) Parse(file string) (*common.SchemaContext, error) {
	schemaPath, err := common.ParseSchemaPath(file)
	if err != nil {
		return nil, err
	}
	packageName := schemaPath.StrippedFileName
	schemaContext := &common.SchemaContext{
		SchemaPath:  *schemaPath,
		PackageName: packageName,
		Classes:     make(map[string]*common.SchemaClass),
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
		//fmt.Printf("%s %s %s\n", triple.Subj, triple.Pred, triple.Obj)

		// Check if the triple describes a class or a property.
		switch triple.Pred.String() {
		case "http://www.w3.org/1999/02/22-rdf-syntax-ns#type":
			if triple.Obj.String() == "http://www.w3.org/2000/01/rdf-schema#Class" {
				schemaContext.Classes[triple.Subj.String()] = &common.SchemaClass{
					Name:    triple.Subj.String(),
					Package: packageName,
				}
			}
		case "http://www.w3.org/1999/02/22-rdf-syntax-ns#about":
			if triple.Subj.String() == "http://www.w3.org/2000/01/rdf-schema#Class" {
				schemaContext.Classes[triple.Obj.String()] = &common.SchemaClass{Name: triple.Obj.String()}
			}
		case "http://www.w3.org/2000/01/rdf-schema#domain", "https://schema.org/domainIncludes":
			if class, ok := schemaContext.Classes[triple.Obj.String()]; ok {
				property := &common.SchemaProperty{
					Name:     triple.Subj.String(),
					Domain:   triple.Obj.String(),
					LangType: "string", // default type when no range is specified
				}
				class.Properties = append(class.Properties, property)
			}
		case "http://www.w3.org/2000/01/rdf-schema#range", "https://schema.org/rangeIncludes":
			for _, class := range schemaContext.Classes {
				for i, property := range class.Properties {
					if property.Name == triple.Subj.String() {
						class.Properties[i].Comment = "//" + triple.Obj.String()
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
							class.Properties[i].LangType = "*" + common.LocalName(triple.Obj.String())
							//class.Properties[i] = nil
						}
					}
				}
			}
		case "<http://www.w3.org/1999/02/22-rdf-syntax-ns#about":

			// Capture OWL properties and classes!!!!!!

			if triple.Subj.String() == "http://www.w3.org/2002/07/owl#Class" {
				schemaContext.Classes[triple.Obj.String()] = &common.SchemaClass{Name: triple.Obj.String()}
			}
		}
	}

	return schemaContext, nil
}
