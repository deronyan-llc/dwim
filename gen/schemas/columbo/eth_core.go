package generated

import (
	"time"
)

// Block is a generated struct representing the http://example.org/ethereum/Block class.
type Block struct {
	Miner            *Address  `json:"miner"`            //http://example.org/ethereum/Address
	MixHash          string    `json:"mixhash"`          //http://www.w3.org/2001/XMLSchema#string
	BaseFeePerGas    *Gwei     `json:"basefeepergas"`    //http://example.org/ethereum/Gwei
	Number           int       `json:"number"`           //http://www.w3.org/2001/XMLSchema#integer
	Hash             string    `json:"hash"`             //http://www.w3.org/2001/XMLSchema#string
	ParentHash       string    `json:"parenthash"`       //http://www.w3.org/2001/XMLSchema#string
	Nonce            string    `json:"nonce"`            //http://www.w3.org/2001/XMLSchema#string
	Sha3Uncles       string    `json:"sha3uncles"`       //http://www.w3.org/2001/XMLSchema#string
	LogsBloom        string    `json:"logsbloom"`        //http://www.w3.org/2001/XMLSchema#string
	TransactionsRoot string    `json:"transactionsroot"` //http://www.w3.org/2001/XMLSchema#string
	StateRoot        string    `json:"stateroot"`        //http://www.w3.org/2001/XMLSchema#string
	ReceiptsRoot     string    `json:"receiptsroot"`     //http://www.w3.org/2001/XMLSchema#string
	GasUsed          int       `json:"gasused"`          //http://www.w3.org/2001/XMLSchema#integer
	GasLimit         int       `json:"gaslimit"`         //http://www.w3.org/2001/XMLSchema#integer
	Difficulty       int       `json:"difficulty"`       //http://www.w3.org/2001/XMLSchema#integer
	Timestamp        time.Time `json:"timestamp"`        //http://www.w3.org/2001/XMLSchema#dateTime
	ExtraData        string    `json:"extradata"`        //http://www.w3.org/2001/XMLSchema#string
}

// Address is a generated struct representing the http://example.org/ethereum/Address class.
type Address struct {
	Address string `json:"address"` //http://www.w3.org/2001/XMLSchema#string
}

// Gwei is a generated struct representing the http://example.org/ethereum/Gwei class.
type Gwei struct {
}
