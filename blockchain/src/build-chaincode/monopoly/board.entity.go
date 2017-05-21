package monopoly

type Board struct {
	Hash            string     `json:"hash"`
	Positions       []Position `json:"positions"`
	Player1         Player     `json:"player1"`
	Player2         Player     `json:"player2"`
	TransactionHash string     `json:"txhash"`
	ChainHeight     int        `json:"chainHeight"`
}
type ChainTX struct {
	Height            int    `json:"height"`
	CurrentBlockHash  string `json:"currentBlockHash"`
	PreviousBlockHash string `json:"previousBlockHash"`
}
