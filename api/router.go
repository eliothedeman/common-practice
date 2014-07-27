package api

import (
	"github.com/eliothedeman/common-practice/search"
	"net/http"
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
}

// Serve servers the api via http
func Serve() {
	http.HandleFunc("/", Route)
	http.ListenAndServe(":8080", nil)
}
