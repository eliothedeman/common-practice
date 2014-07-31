package search

import (
	"strconv"
)

// KeyPointer is the basic struct for storing index pointers in files
type KeyPointer struct {
	Key     string `json:"key"`
	Val     string `json:"val"`
	Pointer int64  `json:"pointer"`
}

// String makes returns string representation of the KeyPointer
func (k *KeyPointer) String() string {
	return "Key: " + k.Key + " Pointer: " + strconv.Itoa(int(k.Pointer))
}
