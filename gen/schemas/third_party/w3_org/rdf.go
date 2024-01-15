package generated

// Seq is a generated struct representing the http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq class.
type Seq struct {
}

// Alt is a generated struct representing the http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt class.
type Alt struct {
}

// List is a generated struct representing the http://www.w3.org/1999/02/22-rdf-syntax-ns#List class.
type List struct {
	First *Resource `json:"first"` //http://www.w3.org/2000/01/rdf-schema#Resource
	Rest  *List     `json:"rest"`  //http://www.w3.org/1999/02/22-rdf-syntax-ns#List
}

// CompoundLiteral is a generated struct representing the http://www.w3.org/1999/02/22-rdf-syntax-ns#CompoundLiteral class.
type CompoundLiteral struct {
	Language  string `json:"language"`
	Direction string `json:"direction"`
}

// Property is a generated struct representing the http://www.w3.org/1999/02/22-rdf-syntax-ns#Property class.
type Property struct {
}

// Statement is a generated struct representing the http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement class.
type Statement struct {
	Subject   *Resource `json:"subject"`   //http://www.w3.org/2000/01/rdf-schema#Resource
	Predicate *Resource `json:"predicate"` //http://www.w3.org/2000/01/rdf-schema#Resource
	Object    *Resource `json:"object"`    //http://www.w3.org/2000/01/rdf-schema#Resource
}

// Bag is a generated struct representing the http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag class.
type Bag struct {
}
