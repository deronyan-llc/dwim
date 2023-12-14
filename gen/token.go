package main

type Token struct {
	TokenId int `json:"tokenid"`
	Owner Agent `json:"owner"`
	ContractAddress string `json:"contractaddress"`
	TokenURI anyURI `json:"tokenuri"`
	Transfer TransferEvent `json:"transfer"`
}
