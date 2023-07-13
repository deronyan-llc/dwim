package main

type Block struct {
	Number           int    `json:"number"`
	Hash             string `json:"hash"`
	ParentHash       string `json:"parenthash"`
	Nonce            string `json:"nonce"`
	Sha3Uncles       string `json:"sha3uncles"`
	LogsBloom        string `json:"logsbloom"`
	TransactionsRoot string `json:"transactionsroot"`
	StateRoot        string `json:"stateroot"`
	ReceiptsRoot     string `json:"receiptsroot"`
	GasUsed          int    `json:"gasused"`
	GasLimit         int    `json:"gaslimit"`
	Difficulty       int    `json:"difficulty"`
}
