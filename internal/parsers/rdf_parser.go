package parsers

import (
	"bytes"
	"os"
	"slices"

	"github.com/deronyan-llc/dwim/internal/common"
	"github.com/deronyan-llc/dwim/internal/consts"
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
		Classes:     make(map[rdf.Term]*common.SchemaClass),
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
		switch triple.Pred {
		case consts.Rdf_type:
			if triple.Obj == consts.Rdfs_Class || triple.Obj == consts.Void_DatasetDescription || triple.Obj == consts.Void_Dataset {
				schemaContext.Classes[triple.Subj] = &common.SchemaClass{
					Term:       triple.Subj,
					TermString: triple.Subj.String(),
					Package:    packageName,
				}
			}
		case consts.Rdf_about:
			if triple.Subj == consts.Rdfs_Class || triple.Obj == consts.Void_DatasetDescription || triple.Obj == consts.Void_Dataset {
				schemaContext.Classes[triple.Obj] = &common.SchemaClass{
					Term:       triple.Obj,
					TermString: triple.Obj.String(),
					Package:    packageName,
				}
			}
		// checking both of these can be deprecated with a proper OWL reasoner
		case consts.Rdfs_domain, consts.SchemaOrg_domainIncludes:
			if class, ok := schemaContext.Classes[triple.Obj]; ok {
				property := &common.SchemaProperty{
					Term:       triple.Subj,
					TermString: triple.Subj.String(),
					Domain:     triple.Obj,
					GoLangType: "string", // default type when no range is specified
				}
				class.Properties = append(class.Properties, property)
			}
		case consts.Rdfs_range, consts.SchemaOrg_rangeIncludes:
			for _, class := range schemaContext.Classes {
				for i, property := range class.Properties {
					if property.Term == triple.Subj {
						class.Properties[i].Comment = "//" + triple.Obj.String()
						class.Properties[i].Range = triple.Obj
						class.Properties[i].GoLangType = triple.Obj.String()
						//classes[triple.Obj.String()] = class

						/*
						 * TODO - make this better, more abstract using a map of RDF types to Go types.
						 * Likely, some interface parent is needed for all the types.
						 */

						// assign the lang-specific data type for the Range
						objStr := triple.Obj.String()
						switch objStr {
						case common.XsdInteger.Str, common.XsdInt.Str: // "http://www.w3.org/2001/XMLSchema#integer":
							class.Properties[i].GoLangType = "int"
						case common.XsdString.Str: // "http://www.w3.org/2001/XMLSchema#string":
							class.Properties[i].GoLangType = "string"
						case common.XsdDateTime.Str: // "http://www.w3.org/2001/XMLSchema#dateTime":
							class.Properties[i].GoLangType = "time.Time"
							if !slices.Contains(class.Imports, "time") {
								class.Imports = append(class.Imports, "time")
							}
						default:
							class.Properties[i].GoLangType = "*" + common.LocalName(triple.Obj.String())
							//class.Properties[i] = nil
						}
					}
				}
			}
		case consts.Rdf_about:
			// Capture OWL properties and classes!!!!!!
			// Should be used via an OWL reasoner
			if triple.Subj == consts.Owl_Class {
				schemaContext.Classes[triple.Obj] = &common.SchemaClass{
					Term:       triple.Obj,
					TermString: triple.Obj.String(),
					Package:    packageName,
				}
			}
		}
	}

	return schemaContext, nil
}
