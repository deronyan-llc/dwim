package generated

// Token is a generated struct representing the https://www.example.com/erc721/Token class.
type Token struct {
	TokenId         int           `json:"tokenid"`
	Owner           Agent         `json:"owner"`
	ContractAddress string        `json:"contractaddress"`
	TokenURI        anyURI        `json:"tokenuri"`
	Transfer        TransferEvent `json:"transfer"`
}

// TransferEvent is a generated struct representing the https://www.example.com/erc721/TransferEvent class.
type TransferEvent struct {
	From string `json:"from"`
	To   string `json:"to"`
}
