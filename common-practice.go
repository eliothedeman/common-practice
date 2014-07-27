package main

import (
	"github.com/eliothedeman/common-practice/api"
)

func main() {
	api.LoadHandlers()
	api.LoadDbs()
	api.Serve()
}
