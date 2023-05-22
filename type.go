package main

type SolStruct struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Context struct {
			Slot int `json:"slot"`
		} `json:"context"`
		Value int `json:"value"`
	} `json:"result"`
}
