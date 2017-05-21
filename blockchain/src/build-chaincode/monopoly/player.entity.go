package monopoly

type Player struct {
	ID              int        `json:"id"`
	Balance         int        `json:"balance"`
	PositionsOwned  []Position `json:"positions"`
	CurrentPosition Position   `json:"currentPosition"`
	CurrentTurn     bool       `json:"currentTurn"`
	HasRolled       bool       `json:"hasRolled"`
}
