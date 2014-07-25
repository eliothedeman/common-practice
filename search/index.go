package search

import (
	"strconv"
	"sync"
)

type KeyPointer struct {
	Key     string
	Pointer int64
}

func (k *KeyPointer) String() string {
	return "Key: " + k.Key + " Pointer: " + strconv.Itoa(int(k.Pointer))
}

// Index is the datatype that will be searched against in the ndb instance
type Index struct {
	Children map[uint64]*KeyPointer
	*sync.Mutex
	ThisKey string
}

// lock the mutex
func (i *Index) lock() {
	i.lock()
}

// unlock the mutex
func (i *Index) unlock() {
	i.unlock()
}

// UpdateIndex adds a new FetchKey to the index
func (i *Index) UpdateIndex(k *KeyPointer) {
	i.lock()
	i.Children[uint64(len(i.Children))] = k
	i.unlock()
}
