package search

// SplitData is a helper function that splits the data string into all of its searchable parts
func SplitData(data string) []*KeyPointer {
	keys := make([]*KeyPointer, ((len(data) * (len(data) + 1)) / 2))
	x := 0
	for l := 0; l < len(data); l++ {
		for i := 0; i+l < len(data); i++ {
			keys[x] = &KeyPointer{Key: data[l : l+i], Pointer: int64(i + l)}

			x += 1

		}
	}
	return keys
}
