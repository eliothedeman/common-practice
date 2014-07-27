package searchAPI

import (
	"github.com/eliothedeman/common-practice/search"
)

var (
	HANDLERS map[string]APIHandler
	DBS      map[string]*search.SearchDb
)

// LoadHandlers loads the handlers into the global map
func LoadHandlers() {
	HANDLERS = make(map[string]APIHandler)
	HANDLERS["insert"] = Insert
}

// LoadDbs loads the SearchDbs in to the global map
func LoadDbs() {
	DBS = make(map[string]*search.SearchDb)
	DBS["pitch"] = search.NewSearchDb("pitch")
	DBS["rhythm"] = search.NewSearchDb("rhythm")
	DBS["key"] = search.NewSearchDb("key")
}
