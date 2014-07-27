package api

import (
	"github.com/eliothedeman/common-practice/api/searchAPI"
	"net/http"
)

// Serve servers the api via http
func Serve(port string) {
	http.HandleFunc("/search", searchAPI.Route)
	http.ListenAndServe(":"+port, nil)
}
