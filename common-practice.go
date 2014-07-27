package main

import (
	"github.com/eliothedeman/common-practice/api"
)

const (
	API_PORT = "8080"
)

func main() {
	api.Serve(API_PORT)
}
