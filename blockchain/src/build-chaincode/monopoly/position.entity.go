package monopoly

type Position struct {
	ID        int    `json:"id"`
	BelongsTo string `json:"belongsTo"`
	Cost      int    `json:"cost"`
	Hotels    int    `json:"hotels"`
	Houses    int    `json:"houses"`
	Type      string `json:"type"`
}
