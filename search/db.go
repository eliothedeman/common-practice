package search

import (
	"errors"
	"github.com/eliothedeman/nbd/core"
)

// SearchDb holds an instance of an nbd used for doing searches
type SearchDb struct {
	Manager *core.Manager
}

// NewSearchDb creates, and returns a new pointer to a SearchDb
func NewSearchDb(name string) *SearchDb {
	s := &SearchDb{}
	m := core.NewManager(1000, name)
	m.DoneChan = make(chan bool)
	s.Manager = m
	<-m.DoneChan
	return s
}

// Insert inserts a piece of data into the nbd instance
func (s *SearchDb) Insert(key string, value interface{}) {
	responseChan := make(chan core.Response)
	s.Manager.RequestChan <- core.Request{Action: 1, Id: key, Data: value, ResponseChan: responseChan, NewRequest: true}
	<-responseChan

}

// Select reads a piece of data from the ndb instance
func (s *SearchDb) Select(key string) interface{} {
	responseChan := make(chan core.Response)
	s.Manager.RequestChan <- core.Request{Action: 0, Id: key, ResponseChan: responseChan}
	resp := <-responseChan
	return resp.Data
}

// SelectAll grabs all children associated with a key
func (s *SearchDb) SelectAll(key string) map[uint64]*KeyPointer {
	elements := s.Select(key)
	switch elements.(type) {
	case *Index:
		return elements.(*Index).Children
	default:
		return nil
	}
}

// UpdateChildren updates the children of an Index
func (s *SearchDb) UpdateChildren(key string, child *KeyPointer) error {
	element := s.Select(key)
	switch element.(type) {
	case *Index:
		element.(*Index).UpdateIndex(child)
		return nil
	default:
		return errors.New("Index not found")
	}
}

// CreateIndexIfNotExist creates an index in the database if it doesn't already exist
func (s *SearchDb) CreateIndexIfNotExist(key string) {
	if s.Select(key) != nil {
		return
	}
	n := NewIndex()
	n.ThisKey = key
	s.Insert(key, n)
}

// CreateAndUpdateIndicies creates and updates a new set of indicies from a dataset
func (s *SearchDb) CreateAndUpdateIndicies(dataSet, tag string) error {
	arr := SplitData(dataSet, tag)
	var err error
	for i := 0; i < len(arr); i++ {
		s.CreateIndexIfNotExist(arr[i].Key)
		err = s.UpdateChildren(arr[i].Key, arr[i])
		if err != nil {
			return err
		}
	}
	return nil
}
