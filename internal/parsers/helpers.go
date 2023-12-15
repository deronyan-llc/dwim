package parsers

import (
	"strings"

	"github.com/knakk/rdf"
)

var FileExtensionsToRdfFormats map[string]rdf.Format = map[string]rdf.Format{
	"ttl": rdf.Turtle,
	"xsd": rdf.RDFXML,
	"xml": rdf.RDFXML,
	"n3":  rdf.NTriples,
	"nq":  rdf.NQuads,
}

// returns the RDF format for the given file extension, if it exists.
// also returns the file name stripped of its extension, and the extension separately
func FileExtToRdfFormat(fileName string) (rdf.Format, string, string) {
	dotPos := strings.LastIndex(fileName, ".")
	if dotPos == -1 {
		return rdf.Format(-1), "", ""
	}
	for ext, format := range FileExtensionsToRdfFormats {
		if strings.HasSuffix(fileName, "."+ext) {
			strippedFileName := strings.TrimSuffix(fileName, "."+ext)
			return format, strippedFileName, ext
		}
	}
	return rdf.Format(-1), "", ""
}
