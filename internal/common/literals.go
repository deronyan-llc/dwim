package common

/*
 * Some of these types are not used in the generated code, but are included for completeness.
 *
 */

// IRI represents a RDF IRI resource.
type IRI struct {
	Str string
}

// The XML schema built-in datatypes (xsd):
// https://dvcs.w3.org/hg/rdf/raw-file/default/rdf-concepts/index.html#xsd-datatypes
var (
	// Core types:                                                    // Corresponding Go datatype:

	XsdString  = IRI{Str: "http://www.w3.org/2001/XMLSchema#string"}  // string
	XsdBoolean = IRI{Str: "http://www.w3.org/2001/XMLSchema#boolean"} // bool
	XsdDecimal = IRI{Str: "http://www.w3.org/2001/XMLSchema#decimal"} // float64
	XsdInteger = IRI{Str: "http://www.w3.org/2001/XMLSchema#integer"} // int

	// IEEE floating-point numbers:
	XsdDouble = IRI{Str: "http://www.w3.org/2001/XMLSchema#double"} // float64
	XsdFloat  = IRI{Str: "http://www.w3.org/2001/XMLSchema#float"}  // float64

	// Time and date:
	XsdDate          = IRI{Str: "http://www.w3.org/2001/XMLSchema#date"}
	XsdTime          = IRI{Str: "http://www.w3.org/2001/XMLSchema#time"}
	XsdDateTime      = IRI{Str: "http://www.w3.org/2001/XMLSchema#dateTime"} // time.Time
	XsdDateTimeStamp = IRI{Str: "http://www.w3.org/2001/XMLSchema#dateTimeStamp"}

	// Recurring and partial dates:
	XsdYear              = IRI{Str: "http://www.w3.org/2001/XMLSchema#gYear"}
	XsdMonth             = IRI{Str: "http://www.w3.org/2001/XMLSchema#gMonth"}
	XsdDay               = IRI{Str: "http://www.w3.org/2001/XMLSchema#gDay"}
	XsdYearMonth         = IRI{Str: "http://www.w3.org/2001/XMLSchema#gYearMonth"}
	XsdDuration          = IRI{Str: "http://www.w3.org/2001/XMLSchema#Duration"}
	XsdYearMonthDuration = IRI{Str: "http://www.w3.org/2001/XMLSchema#yearMonthDuration"}
	XsdDayTimeDuration   = IRI{Str: "http://www.w3.org/2001/XMLSchema#dayTimeDuration"}

	// Limited-range integer numbers
	XsdByte  = IRI{Str: "http://www.w3.org/2001/XMLSchema#byte"}  // []byte
	XsdShort = IRI{Str: "http://www.w3.org/2001/XMLSchema#short"} // int16
	XsdInt   = IRI{Str: "http://www.w3.org/2001/XMLSchema#int"}   // int32
	XsdLong  = IRI{Str: "http://www.w3.org/2001/XMLSchema#long"}  // int64

	// Various
	RdfLangString = IRI{Str: "http://www.w3.org/1999/02/22-rdf-syntax-ns#langString"} // string
	XmlLiteral    = IRI{Str: "http://www.w3.org/1999/02/22-rdf-syntax-ns#XMLLiteral"} // string
)
