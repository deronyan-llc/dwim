package consts

import "github.com/deronyan-llc/rdf/rdf"

func mustNewIRI(str string) rdf.Term {
	t, err := rdf.NewIRI(str)
	if err != nil {
		panic(err)
	}
	return t
}

var (
	// RDFS Terms
	Rdfs_Class         = mustNewIRI("http://www.w3.org/2000/01/rdf-schema#Class")
	Rdfs_domain        = mustNewIRI("http://www.w3.org/2000/01/rdf-schema#domain")
	Rdfs_range         = mustNewIRI("http://www.w3.org/2000/01/rdf-schema#range")
	Rdfs_subClassOf    = mustNewIRI("http://www.w3.org/2000/01/rdf-schema#subClassOf")
	Rdfs_subPropertyOf = mustNewIRI("http://www.w3.org/2000/01/rdf-schema#subPropertyOf")
	Rdfs_label         = mustNewIRI("http://www.w3.org/2000/01/rdf-schema#label")
	Rdfs_comment       = mustNewIRI("http://www.w3.org/2000/01/rdf-schema#comment")

	// RDF Terms
	Rdf_type  = mustNewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	Rdf_about = mustNewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#about")

	// schema.org Terms
	SchemaOrg_domainIncludes = mustNewIRI("https://schema.org/domainIncludes")
	SchemaOrg_rangeIncludes  = mustNewIRI("https://schema.org/rangeIncludes")

	// OWL Terms
	Owl_Class = mustNewIRI("http://www.w3.org/2002/07/owl#Class")

	// VOID Terms
	Void_DatasetDescription = mustNewIRI("http://rdfs.org/ns/void#DatasetDescription")
	Void_Dataset            = mustNewIRI("http://rdfs.org/ns/void#Dataset")
)
