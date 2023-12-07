package main

import (
	"strings"

	"github.com/knakk/rdf"
)

type Class struct {
	Name       string
	Properties []Property
}

type Property struct {
	Name string
	Type string
}

func main() {
	schema := `
	@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .
@prefix schema: <http://schema.org/> .
@prefix dcterms: <http://purl.org/dc/terms/> .
@prefix foaf: <http://xmlns.com/foaf/0.1/> .
@prefix erc721: <https://www.example.com/erc721/> .

erc721:Token
    a rdfs:Class ;
    rdfs:label "Non-fungible Token" ;
    rdfs:comment "A non-fungible token (NFT) as specified by ERC-721" ;
    rdfs:subClassOf schema:Product .

erc721:tokenId
    a rdf:Property ;
    rdfs:label "Token ID" ;
    rdfs:comment "The unique identifier of the token" ;
    rdfs:domain erc721:Token ;
    rdfs:range xsd:integer .

erc721:owner
    a rdf:Property ;
    rdfs:label "Token owner" ;
    rdfs:comment "The owner of the token" ;
    rdfs:domain erc721:Token ;
    rdfs:range foaf:Agent .

erc721:contractAddress
    a rdf:Property ;
    rdfs:label "Contract address" ;
    rdfs:comment "The address of the ERC-721 smart contract" ;
    rdfs:domain erc721:Token ;
    rdfs:range xsd:string .

erc721:tokenURI
    a rdf:Property ;
    rdfs:label "Token URI" ;
    rdfs:comment "The URI containing metadata about the token" ;
    rdfs:domain erc721:Token ;
    rdfs:range xsd:anyURI .

erc721:transfer
    a rdf:Property ;
    rdfs:label "Transfer event" ;
    rdfs:comment "An event emitted when a token is transferred" ;
    rdfs:domain erc721:Token ;
    rdfs:range erc721:TransferEvent .

erc721:TransferEvent
    a rdfs:Class ;
    rdfs:label "Transfer event" ;
    rdfs:comment "An event emitted when a token is transferred" ;
    rdfs:subClassOf schema:Event .

erc721:from
    a rdf:Property ;
    rdfs:label "From address" ;
    rdfs:comment "The address transferring the token" ;
    rdfs:domain erc721:TransferEvent ;
    rdfs:range xsd:string .

erc721:to
    a rdf:Property ;
    rdfs:label "To address" ;
    rdfs:comment "The address receiving the token" ;
    rdfs:domain erc721:TransferEvent ;
    rdfs:range xsd:string .
`
	parseRDFSchema(schema)
}
func parseRDFSchema(schema string) ([]Class, error) {
	var classes []Class

	decoder := rdf.NewQuadDecoder(strings.NewReader(schema), rdf.Turtle)
	g, err := decoder.DecodeAll()
	if err != nil {
		return nil, err
	}
	for _, quad := range g {
		if quad.Pred.String() != "type" || quad.Obj.String() != "Class" {
			/*
				cls := Class{Name: class.Frag()}
				for _, prop := range g.SubjectsOf(rdf.RDF("type"), rdf.OWL("ObjectProperty")) {
					if r, ok := g.One(prop, rdf.RDFS("domain")).(rdf.NamedNode); ok && r == class {
						cls.Properties = append(cls.Properties, Property{Name: prop.Frag()})
					}
				}
				classes = append(classes, cls)
			*/
		}
	}

	return classes, nil

}

func goStruct(class Class) string {
	var structBuilder strings.Builder
	structBuilder.WriteString("type ")
	structBuilder.WriteString(class.Name)
	structBuilder.WriteString(" struct {\n")

	for _, prop := range class.Properties {
		structBuilder.WriteString("\t")
		structBuilder.WriteString(prop.Name)
		structBuilder.WriteString(" ")
		structBuilder.WriteString(prop.Type)
		structBuilder.WriteString(" `json:\"")
		structBuilder.WriteString(strings.ToLower(prop.Name))
		structBuilder.WriteString("\" bson:\"")
		structBuilder.WriteString(strings.ToLower(prop.Name))
		structBuilder.WriteString("\"`\n")
	}

	structBuilder.WriteString("}\n\n")
	return structBuilder.String()
}

func generateGoSource(classes []Class) string {
	var sourceBuilder strings.Builder

	sourceBuilder.WriteString("package generated\n\n")

	for _, class := range classes {
		sourceBuilder.WriteString(goStruct(class))
	}

	return sourceBuilder.String()
}
