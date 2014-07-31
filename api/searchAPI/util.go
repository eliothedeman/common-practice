package searchAPI

func DumpString(d []uint8) string {

	for i := 0; i < len(d); i++ {
		d[i] = d[i] + 64
	}
	return string(d)
}
