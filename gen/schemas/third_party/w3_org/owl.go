package generated

// AnnotationProperty is a generated struct representing the http://www.w3.org/2002/07/owl#AnnotationProperty class.
type AnnotationProperty struct {
}

// Axiom is a generated struct representing the http://www.w3.org/2002/07/owl#Axiom class.
type Axiom struct {
}

// DeprecatedClass is a generated struct representing the http://www.w3.org/2002/07/owl#DeprecatedClass class.
type DeprecatedClass struct {
}

// FunctionalProperty is a generated struct representing the http://www.w3.org/2002/07/owl#FunctionalProperty class.
type FunctionalProperty struct {
}

// InverseFunctionalProperty is a generated struct representing the http://www.w3.org/2002/07/owl#InverseFunctionalProperty class.
type InverseFunctionalProperty struct {
}

// IrreflexiveProperty is a generated struct representing the http://www.w3.org/2002/07/owl#IrreflexiveProperty class.
type IrreflexiveProperty struct {
}

// ObjectProperty is a generated struct representing the http://www.w3.org/2002/07/owl#ObjectProperty class.
type ObjectProperty struct {
	InverseOf          *ObjectProperty `json:"inverseof"`          //http://www.w3.org/2002/07/owl#ObjectProperty
	PropertyChainAxiom *List           `json:"propertychainaxiom"` //http://www.w3.org/1999/02/22-rdf-syntax-ns#List
}

// AllDisjointClasses is a generated struct representing the http://www.w3.org/2002/07/owl#AllDisjointClasses class.
type AllDisjointClasses struct {
}

// Annotation is a generated struct representing the http://www.w3.org/2002/07/owl#Annotation class.
type Annotation struct {
}

// Class is a generated struct representing the http://www.w3.org/2002/07/owl#Class class.
type Class struct {
	ComplementOf    *Class `json:"complementof"`    //http://www.w3.org/2002/07/owl#Class
	DisjointUnionOf *List  `json:"disjointunionof"` //http://www.w3.org/1999/02/22-rdf-syntax-ns#List
	DisjointWith    *Class `json:"disjointwith"`    //http://www.w3.org/2002/07/owl#Class
	HasKey          *List  `json:"haskey"`          //http://www.w3.org/1999/02/22-rdf-syntax-ns#List
}

// DeprecatedProperty is a generated struct representing the http://www.w3.org/2002/07/owl#DeprecatedProperty class.
type DeprecatedProperty struct {
}

// NegativePropertyAssertion is a generated struct representing the http://www.w3.org/2002/07/owl#NegativePropertyAssertion class.
type NegativePropertyAssertion struct {
	AssertionProperty *Property `json:"assertionproperty"` //http://www.w3.org/1999/02/22-rdf-syntax-ns#Property
	SourceIndividual  *Thing    `json:"sourceindividual"`  //http://www.w3.org/2002/07/owl#Thing
	TargetIndividual  *Thing    `json:"targetindividual"`  //http://www.w3.org/2002/07/owl#Thing
	TargetValue       *Literal  `json:"targetvalue"`       //http://www.w3.org/2000/01/rdf-schema#Literal
}

// Restriction is a generated struct representing the http://www.w3.org/2002/07/owl#Restriction class.
type Restriction struct {
	AllValuesFrom           *Class              `json:"allvaluesfrom"`           //http://www.w3.org/2000/01/rdf-schema#Class
	Cardinality             *nonNegativeInteger `json:"cardinality"`             //http://www.w3.org/2001/XMLSchema#nonNegativeInteger
	HasSelf                 *Resource           `json:"hasself"`                 //http://www.w3.org/2000/01/rdf-schema#Resource
	HasValue                *Resource           `json:"hasvalue"`                //http://www.w3.org/2000/01/rdf-schema#Resource
	MaxCardinality          *nonNegativeInteger `json:"maxcardinality"`          //http://www.w3.org/2001/XMLSchema#nonNegativeInteger
	MaxQualifiedCardinality *nonNegativeInteger `json:"maxqualifiedcardinality"` //http://www.w3.org/2001/XMLSchema#nonNegativeInteger
	MinCardinality          *nonNegativeInteger `json:"mincardinality"`          //http://www.w3.org/2001/XMLSchema#nonNegativeInteger
	MinQualifiedCardinality *nonNegativeInteger `json:"minqualifiedcardinality"` //http://www.w3.org/2001/XMLSchema#nonNegativeInteger
	OnClass                 *Class              `json:"onclass"`                 //http://www.w3.org/2002/07/owl#Class
	OnDataRange             *Datatype           `json:"ondatarange"`             //http://www.w3.org/2000/01/rdf-schema#Datatype
	OnProperties            *List               `json:"onproperties"`            //http://www.w3.org/1999/02/22-rdf-syntax-ns#List
	OnProperty              *Property           `json:"onproperty"`              //http://www.w3.org/1999/02/22-rdf-syntax-ns#Property
	QualifiedCardinality    *nonNegativeInteger `json:"qualifiedcardinality"`    //http://www.w3.org/2001/XMLSchema#nonNegativeInteger
	SomeValuesFrom          *Class              `json:"somevaluesfrom"`          //http://www.w3.org/2000/01/rdf-schema#Class
}

// DataRange is a generated struct representing the http://www.w3.org/2002/07/owl#DataRange class.
type DataRange struct {
}

// Ontology is a generated struct representing the http://www.w3.org/2002/07/owl#Ontology class.
type Ontology struct {
	BackwardCompatibleWith *Ontology `json:"backwardcompatiblewith"` //http://www.w3.org/2002/07/owl#Ontology
	Imports                *Ontology `json:"imports"`                //http://www.w3.org/2002/07/owl#Ontology
	IncompatibleWith       *Ontology `json:"incompatiblewith"`       //http://www.w3.org/2002/07/owl#Ontology
	PriorVersion           *Ontology `json:"priorversion"`           //http://www.w3.org/2002/07/owl#Ontology
	VersionIRI             *Ontology `json:"versioniri"`             //http://www.w3.org/2002/07/owl#Ontology
}

// ReflexiveProperty is a generated struct representing the http://www.w3.org/2002/07/owl#ReflexiveProperty class.
type ReflexiveProperty struct {
}

// TransitiveProperty is a generated struct representing the http://www.w3.org/2002/07/owl#TransitiveProperty class.
type TransitiveProperty struct {
}

// AllDifferent is a generated struct representing the http://www.w3.org/2002/07/owl#AllDifferent class.
type AllDifferent struct {
	DistinctMembers *List `json:"distinctmembers"` //http://www.w3.org/1999/02/22-rdf-syntax-ns#List
}

// AllDisjointProperties is a generated struct representing the http://www.w3.org/2002/07/owl#AllDisjointProperties class.
type AllDisjointProperties struct {
}

// AsymmetricProperty is a generated struct representing the http://www.w3.org/2002/07/owl#AsymmetricProperty class.
type AsymmetricProperty struct {
}

// DatatypeProperty is a generated struct representing the http://www.w3.org/2002/07/owl#DatatypeProperty class.
type DatatypeProperty struct {
}

// NamedIndividual is a generated struct representing the http://www.w3.org/2002/07/owl#NamedIndividual class.
type NamedIndividual struct {
}

// OntologyProperty is a generated struct representing the http://www.w3.org/2002/07/owl#OntologyProperty class.
type OntologyProperty struct {
}

// SymmetricProperty is a generated struct representing the http://www.w3.org/2002/07/owl#SymmetricProperty class.
type SymmetricProperty struct {
}
