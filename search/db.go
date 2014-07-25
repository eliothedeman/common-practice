package search

import (
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

	go s.Manager.Manage()
	<-m.DoneChan
	return s
}

// Insert inserts a piece of data into the nbd instance
func (s *SearchDb) Insert(key string, value interface{}) {
	responseChan := make(chan core.Response)
	core.RequestChan <- core.Request{Action: 1, Id: key, Data: value, ResponseChan: responseChan, NewRequest: true}
	<-responseChan

}

// Select reads a piece of data from the ndb instance
func (s *SearchDb) Select(key string) interface{} {
	responseChan := make(chan core.Response)
	core.RequestChan <- core.Request{Action: 0, Id: key, ResponseChan: responseChan}
	resp := <-responseChan
	return resp.Data
}
