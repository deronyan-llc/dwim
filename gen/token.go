package main

type Token struct {
	TokenId int `json:"tokenid"`
	Owner http://xmlns.com/foaf/0.1/Agent `json:"owner"`
	ContractAddress string `json:"contractaddress"`
	TokenURI http://www.w3.org/2001/XMLSchema#anyURI `json:"tokenuri"`
}
