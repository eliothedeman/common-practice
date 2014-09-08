package api

import (
	"net/http"
)

// Handler interface describes a custom handler for a pre-parsed request
type Handler func(http.ResponseWriter, *Request)
