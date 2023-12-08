package main

type Block struct {
	Miner http://example.org/ethereum/Address `json:"miner"`
	MixHash string `json:"mixhash"`
	BaseFeePerGas http://example.org/ethereum/Gwei `json:"basefeepergas"`
	Number int `json:"number"`
	Hash string `json:"hash"`
	ParentHash string `json:"parenthash"`
	Nonce string `json:"nonce"`
	Sha3Uncles string `json:"sha3uncles"`
	LogsBloom string `json:"logsbloom"`
	TransactionsRoot string `json:"transactionsroot"`
	StateRoot string `json:"stateroot"`
	ReceiptsRoot string `json:"receiptsroot"`
	GasUsed int `json:"gasused"`
	GasLimit int `json:"gaslimit"`
	Difficulty int `json:"difficulty"`
	Timestamp time.Time `json:"timestamp"`
}
