package core

// Piece represents a common practie piece of music
type Piece struct {
	Title string `json:"title"`
	Parts []Part `json:"part"`
}

// Part represents an individuale part in a piece of music i.e. violin 1, trumpet 3
type Part struct {
	Name     string    `json:"name"`
	Measures []Measure `json:"measures"`
}

// Measure represents a measure of music
type Measure struct {
	Notes []Note `json:"notes"`
	Key   Key    `json:"key"`
}

// Note represents a note in a measure of music
type Note struct {
	Pitch Pitch `json:"pitch"`
	Rest  bool  `json:"rest"`
}

// Movement represents a movement from a piece of music
type Movement struct {
	Index      uint8       `json:"index"`
	KeyChanges []KeyChange `json:"keyChanges"`
	Measures   []Measure   `json:"measures"`
}
