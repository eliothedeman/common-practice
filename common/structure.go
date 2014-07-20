package common

// Piece represents a common practie piece of music
type Piece struct {
	Title     string     `json:"title"`
	Movements []Movement `json:"movements"`
}

// Measure represents a measure of music
// Movement represents a movement from a piece of music
type Movement struct {
	Index      uint8       `json:"index"`
	KeyChanges []KeyChange `json:"keyChanges"`
	Measures   []Measure   `json:"measures"`
}
