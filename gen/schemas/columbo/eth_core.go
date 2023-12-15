package generated

import (
	"time"
)

// Block is a generated struct representing the http://example.org/ethereum/Block class.
type Block struct {
	Miner            Address   `json:"miner"`
	MixHash          string    `json:"mixhash"`
	BaseFeePerGas    Gwei      `json:"basefeepergas"`
	Number           int       `json:"number"`
	Hash             string    `json:"hash"`
	ParentHash       string    `json:"parenthash"`
	Nonce            string    `json:"nonce"`
	Sha3Uncles       string    `json:"sha3uncles"`
	LogsBloom        string    `json:"logsbloom"`
	TransactionsRoot string    `json:"transactionsroot"`
	StateRoot        string    `json:"stateroot"`
	ReceiptsRoot     string    `json:"receiptsroot"`
	GasUsed          int       `json:"gasused"`
	GasLimit         int       `json:"gaslimit"`
	Difficulty       int       `json:"difficulty"`
	Timestamp        time.Time `json:"timestamp"`
	ExtraData        string    `json:"extradata"`
}

// Address is a generated struct representing the http://example.org/ethereum/Address class.
type Address struct {
	Address string `json:"address"`
}

// Gwei is a generated struct representing the http://example.org/ethereum/Gwei class.
type Gwei struct {
}
