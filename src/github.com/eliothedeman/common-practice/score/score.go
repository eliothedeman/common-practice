package score

import "github.com/eliothedeman/go-mxl"

//go:generate msgp $GOFILE

// Score represents a piece of music
type Score struct {
	Composer  string     `msg:"composer"`
	Licence   string     `msg:"licence"`
	Title     string     `msg:"title"`
	Movements []Movement `msg:"movements"`
}

// LoadFromMXL takes a parsed mxl file, and loads it's data into a score
func LoadFromMXL(m *mxl.MXLDoc) *Score {
	s := &Score{}
	s.Composer = m.Composer
	s.Licence = m.Rights
	s.Title = m.Identification.Title
	return s
}

// Movement represents a movement in a piece of music
type Movement struct {
	Name  string `msg:"name"`
	Index int    `msg:"index"`
	Parts []Part `msg:"parts"`
}

type Part struct {
	Insturment string    `msg:"insturment"`
	Measures   []Measure `msg:"measures"`
}

// TimeSignature represents the time signature of a measure of music
type TimeSignature struct {
	Beat  uint8 `msg:"beat"`  // top number
	Value uint8 `msg:"value"` // bottom number
}

// Measure represents a measure of music
type Measure struct {
	TimeSignature `msg:"time_signature"`
	Index         uint32 `msg:"index"`
	Notes         []Note `msg:"notes"`
}

// Note represents a note of music
type Note struct {
	Step      uint8  `msg:"step"`
	Octave    uint8  `msg:"octave"`
	Duration  uint16 `msg:"duration"`
	Placement uint16 `msg:"placement"`
}
