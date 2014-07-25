package search

import (
	"testing"
)

func TestSplitData(t *testing.T) {
	d := "asdfgasdfasdkfjasldkjfal;sdjf;alsdkf"
	SplitData(d)
}

func TestInsert(t *testing.T) {
	s := NewSearchDb("testing")
	s.Insert("test", 1)
}
