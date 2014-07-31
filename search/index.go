package search

import (
	"encoding/json"
	"sync"
)

// Index is the datatype that will be searched against in the ndb instance
type Index struct {
	Children map[uint64]*KeyPointer
	*sync.Mutex
	ThisKey string
}

// lock the mutex
func (i *Index) lock() {
	i.Mutex.Lock()
}

// unlock the mutex
func (i *Index) unlock() {
	i.Mutex.Unlock()
}

// UpdateIndex adds a new FetchKey to the index
func (i *Index) UpdateIndex(k *KeyPointer) {
	i.lock()
	i.Children[uint64(len(i.Children))] = k
	i.unlock()
}

// NewIndex creates and returns a new pointer to an Index
func NewIndex() *Index {
	i := &Index{Mutex: &sync.Mutex{}}
	i.Children = make(map[uint64]*KeyPointer)
	return i
}

// MarshalJSON  returns a json encoded version of the Index
func (i *Index) MarshalJSON() ([]byte, error) {
	// convert the map of children to an array
	arr := make([]*KeyPointer, len(i.Children))
	for x := range arr {
		arr[x] = i.Children[uint64(x)]
	}
	// create a map for encoding
	m := make(map[string]interface{})
	m["children"] = arr
	m["key"] = i.ThisKey

	return json.Marshal(m)
}
