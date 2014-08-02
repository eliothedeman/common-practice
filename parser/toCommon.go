package parser

import (
	"github.com/eliothedeman/common-practice/core"
	"github.com/eliothedeman/go-mxl"
)

var (
	PITCH_TO_INT = map[string]int{
		"a": 0,
		"b": 2,
		"c": 3,
		"d": 5,
		"e": 7,
		"f": 8,
		"g": 10,
		"A": 0,
		"B": 2,
		"C": 3,
		"D": 5,
		"E": 7,
		"F": 8,
		"G": 10,
	}
)

// toCommon converts a parsed mxl doc to a common practice score
type toCommon struct {
	Doc *mxl.MXLDoc
}

// Do converts an mxl piece to a core.Piece
func (t *toCommon) Do(m *Manager) {
	p := &core.Piece{}
	p.Title = t.Doc.Title
	p.Parts = make([]core.Part, len(t.Doc.Parts))
	for i := 0; i < len(p.Parts); i++ {

	}

}

func parseParks(p []mxl.Part) []core.Part {
	n := make([]core.Part, len(p))
	for i := range p {
		n[i].Name = p[i].Id
		n[i].Movements[i] = parseMovements(p[i].Measures)
	}
}

func parseMeasures(m []mxl.Measure) []core.Measure {
	n := make([]core.Measure, len(m))
	for i := range m {
		n[i].Key.Fifths = m[i].Atters.Key.Fifths
		n[i].Notes = parseNotes(m[i].Notes)
	}
}

func parseNotes(notes []mxl.Note) []core.Note {
	n := make([]core.Note, len(notes))
	for i := range m {
		n[i].Pitch = core.Pitch(parsePitch(notes[i].Pitch) + notes[i].Accidental)
	}
}

func parsePitch(m mxl.Pitch) core.Pitch {
	return core.Pitch(m.Octave*11 + PITCH_TO_INT[m.Step])
}
