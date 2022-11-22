package models

type Transactions struct {
	Id        uint
	Amount    uint
	Timestamp string
	City      string
	// CreatedAt time.Time
}

type DbResult struct {
	Sum   uint
	Avg   float32
	Count uint
	Min   uint
	Max   uint
}
