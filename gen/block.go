package main

type Block struct {
	Miner Unknown `json:"miner"`
	MixHash Unknown `json:"mixhash"`
	BaseFeePerGas Unknown `json:"basefeepergas"`
	Number Unknown `json:"number"`
	Hash Unknown `json:"hash"`
	ParentHash Unknown `json:"parenthash"`
	Nonce Unknown `json:"nonce"`
	Sha3Uncles Unknown `json:"sha3uncles"`
	LogsBloom Unknown `json:"logsbloom"`
	TransactionsRoot Unknown `json:"transactionsroot"`
	StateRoot Unknown `json:"stateroot"`
	ReceiptsRoot Unknown `json:"receiptsroot"`
	GasUsed Unknown `json:"gasused"`
	GasLimit Unknown `json:"gaslimit"`
	Difficulty Unknown `json:"difficulty"`
	Timestamp Unknown `json:"timestamp"`
}
