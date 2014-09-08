package storage

// Providor represents a storange interface. Storage providors are responsible
// for managing their own connection pools and variable recycling
type Providor interface {
	Get(
		string, // the query string
	) ([]byte, error) // response from the providor
	Set(
		string, // key to store at
		[]byte, // data to store
	) error
}
