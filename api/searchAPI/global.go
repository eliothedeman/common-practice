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
	// setup SearchDbs
	DBS = make(map[string]*search.SearchDb)
	DBS["pitch"] = search.NewSearchDb("pitch")
	DBS["rhythm"] = search.NewSearchDb("rhythm")
	DBS["key"] = search.NewSearchDb("key")
}
