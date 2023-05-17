package main

type Token struct {
	TokenId Integer `json:"tokenid"`
	Owner Agent `json:"owner"`
	ContractAddress String `json:"contractaddress"`
	TokenURI AnyURI `json:"tokenuri"`
}
