package core

var (
	// these represent the pitch alterations from a major scale of each of the modes
	// This should be treated as a constant and shot not be changed.
	MODES = map[string]Mode{
		"major":      []int8{0, 0, 0, 0, 0, 0, 0},
		"dorian":     []int8{0, 0, -1, 0, 0, 0, -1},
		"phrygian":   []int8{0, -1, -1, 0, 0, -1, -1},
		"lydian":     []int8{0, 0, 0, 1, 0, 0, 0},
		"mixolydian": []int8{0, 0, 0, 0, 0, 0, -1},
		"minor":      []int8{0, 0, -1, 0, 0, -1, -1},
		"lorican":    []int8{0, -1, -1, 0, -1, -1, -1},
	}
)

// Mode holds pitch alterations for a given mode
type Mode []int8

// Key holds key signature information
type Key struct {
	// tonal center
	Center int8 `json:"center"`
	// positive numbers are sharps, negative are flats
	Fifths int8 `json:"fifths"`
	// 1-7 represent the 7 church modes
	/*
		1 Ionian(major)
		2 Dorian
		3 Phrygian
		4 Lydian
		5 Mixolydian
		6 Aeolian(minor)
		7 Locrian
	*/
	Mode Mode `json:"mode"`
}

// KeyChange describes a change between two keys
type KeyChange struct {
	OldKey Key `json:"oldKey"`
	NewKey Key `json:"newKey"`
}
