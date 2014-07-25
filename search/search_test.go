package search

import (
	"testing"
)

var (
	s *SearchDb
)

func TestSplitData(t *testing.T) {
	d := "asdfgasdfasdkfjasldkjfal;sdjf;alsdkf"
	SplitData(d)
}

func TestInsert(t *testing.T) {
	s = NewSearchDb("testing")
	s.Insert("test", 1)
}

func TestSelect(t *testing.T) {
	expected := 1
	if expected != s.Select("test") {
		t.Fail()
	}
}
