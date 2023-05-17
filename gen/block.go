package main

type Block struct {
	Number Integer `json:"number"`
	Hash String `json:"hash"`
	ParentHash String `json:"parenthash"`
	Nonce String `json:"nonce"`
	Sha3Uncles String `json:"sha3uncles"`
	LogsBloom String `json:"logsbloom"`
	TransactionsRoot String `json:"transactionsroot"`
	StateRoot String `json:"stateroot"`
	ReceiptsRoot String `json:"receiptsroot"`
	GasUsed Integer `json:"gasused"`
	GasLimit Integer `json:"gaslimit"`
	Difficulty Integer `json:"difficulty"`
}
