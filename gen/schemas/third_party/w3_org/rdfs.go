package generated

// Literal is a generated struct representing the http://www.w3.org/2000/01/rdf-schema#Literal class.
type Literal struct {
}

// Container is a generated struct representing the http://www.w3.org/2000/01/rdf-schema#Container class.
type Container struct {
}

// ContainerMembershipProperty is a generated struct representing the http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty class.
type ContainerMembershipProperty struct {
}

// Datatype is a generated struct representing the http://www.w3.org/2000/01/rdf-schema#Datatype class.
type Datatype struct {
}

// Resource is a generated struct representing the http://www.w3.org/2000/01/rdf-schema#Resource class.
type Resource struct {
	Comment     *Literal  `json:"comment"` //http://www.w3.org/2000/01/rdf-schema#Literal
	Label       *Literal  `json:"label"`   //http://www.w3.org/2000/01/rdf-schema#Literal
	SeeAlso     string    `json:"seealso"`
	IsDefinedBy string    `json:"isdefinedby"`
	Member      *Resource `json:"member"` //http://www.w3.org/2000/01/rdf-schema#Resource
}

// Class is a generated struct representing the http://www.w3.org/2000/01/rdf-schema#Class class.
type Class struct {
	SubClassOf string `json:"subclassof"`
}
