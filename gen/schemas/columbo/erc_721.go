package generated

// Token is a generated struct representing the https://www.example.com/erc721/Token class.
type Token struct {
	TokenId         int            `json:"tokenid"`         //http://www.w3.org/2001/XMLSchema#integer
	Owner           *Agent         `json:"owner"`           //http://xmlns.com/foaf/0.1/Agent
	ContractAddress string         `json:"contractaddress"` //http://www.w3.org/2001/XMLSchema#string
	TokenURI        *anyURI        `json:"tokenuri"`        //http://www.w3.org/2001/XMLSchema#anyURI
	Transfer        *TransferEvent `json:"transfer"`        //https://www.example.com/erc721/TransferEvent
}

// TransferEvent is a generated struct representing the https://www.example.com/erc721/TransferEvent class.
type TransferEvent struct {
	From string `json:"from"` //http://www.w3.org/2001/XMLSchema#string
	To   string `json:"to"`   //http://www.w3.org/2001/XMLSchema#string
}
