package generated

// Linkset is a generated struct representing the http://rdfs.org/ns/void#Linkset class.
type Linkset struct {
	LinkPredicate  *Property `json:"linkpredicate"`  //http://www.w3.org/1999/02/22-rdf-syntax-ns#Property
	ObjectsTarget  *Dataset  `json:"objectstarget"`  //http://rdfs.org/ns/void#Dataset
	SubjectsTarget *Dataset  `json:"subjectstarget"` //http://rdfs.org/ns/void#Dataset
	Target         *Dataset  `json:"target"`         //http://rdfs.org/ns/void#Dataset
}

// TechnicalFeature is a generated struct representing the http://rdfs.org/ns/void#TechnicalFeature class.
type TechnicalFeature struct {
}

// Dataset is a generated struct representing the http://rdfs.org/ns/void#Dataset class.
type Dataset struct {
	Class                 *Class            `json:"class"`                 //http://www.w3.org/2000/01/rdf-schema#Class
	ClassPartition        *Dataset          `json:"classpartition"`        //http://rdfs.org/ns/void#Dataset
	Classes               int               `json:"classes"`               //http://www.w3.org/2001/XMLSchema#integer
	DataDump              *Resource         `json:"datadump"`              //http://www.w3.org/2000/01/rdf-schema#Resource
	DistinctObjects       int               `json:"distinctobjects"`       //http://www.w3.org/2001/XMLSchema#integer
	DistinctSubjects      int               `json:"distinctsubjects"`      //http://www.w3.org/2001/XMLSchema#integer
	Documents             int               `json:"documents"`             //http://www.w3.org/2001/XMLSchema#integer
	Entities              int               `json:"entities"`              //http://www.w3.org/2001/XMLSchema#integer
	ExampleResource       *Resource         `json:"exampleresource"`       //http://www.w3.org/2000/01/rdf-schema#Resource
	Feature               *TechnicalFeature `json:"feature"`               //http://rdfs.org/ns/void#TechnicalFeature
	OpenSearchDescription *Document         `json:"opensearchdescription"` //http://xmlns.com/foaf/0.1/Document
	Properties            int               `json:"properties"`            //http://www.w3.org/2001/XMLSchema#integer
	Property              *Property         `json:"property"`              //http://www.w3.org/1999/02/22-rdf-syntax-ns#Property
	PropertyPartition     *Dataset          `json:"propertypartition"`     //http://rdfs.org/ns/void#Dataset
	RootResource          string            `json:"rootresource"`
	SparqlEndpoint        string            `json:"sparqlendpoint"`
	Subset                *Dataset          `json:"subset"`  //http://rdfs.org/ns/void#Dataset
	Triples               int               `json:"triples"` //http://www.w3.org/2001/XMLSchema#integer
	UriLookupEndpoint     string            `json:"urilookupendpoint"`
	UriRegexPattern       string            `json:"uriregexpattern"`
	UriSpace              *Literal          `json:"urispace"` //http://www.w3.org/2000/01/rdf-schema#Literal
	Vocabulary            string            `json:"vocabulary"`
}

// DatasetDescription is a generated struct representing the http://rdfs.org/ns/void#DatasetDescription class.
type DatasetDescription struct {
}
