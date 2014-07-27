package main

import (
	"github.com/eliothedeman/common-practice/api"
	"github.com/eliothedeman/common-practice/api/searchAPI"
)

const (
	API_PORT = "8080"
)

func main() {
	searchAPI.LoadHandlers()
	searchAPI.LoadDbs()
	api.Serve(API_PORT)
}
