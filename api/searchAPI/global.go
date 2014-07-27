package searchAPI

import (
	"github.com/eliothedeman/common-practice/search"
)

var (
	HANDLERS map[string]APIHandler
	DBS      map[string]*search.SearchDb
)

// setup global vars
func init() {
	// set up handlers
	HANDLERS = make(map[string]APIHandler)
	HANDLERS["insert"] = Insert
	HANDLERS["select"] = Select
	// setup SearchDbs
	DBS = make(map[string]*search.SearchDb)
	// Absolute
	DBS["a-pitch"] = search.NewSearchDb("a-pitch")
	DBS["a-rhythm"] = search.NewSearchDb("a-rhythm")
	DBS["rhythm"] = search.NewSearchDb("rhythm")
	DBS["key"] = search.NewSearchDb("key")
}
