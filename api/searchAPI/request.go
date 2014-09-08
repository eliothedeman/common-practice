package searchAPI

import (
	"errors"
	"net/http"
)

// Request represents an api request
type Request struct {
	Action          string  `json:"action"`
	Data            string  `json:"data"`
	Version         float32 `json:"version"`
	Target          string  `json:"target"`
	Location        string  `json:"location"`
	Response        map[string]interface{}
	ResponseWritter http.ResponseWriter
}

// Validate is a sanity check for a request. If request is valid, return nil
func (r *Request) Validate() error {

	// check that the target and action are not empty strings
	if r.Action == "" || r.Target == "" {
		return errors.New("target and action cannot be empty strings")
	}

	// make sure there is a handler for the requsts action
	if _, ok := HANDLERS[r.Action]; !ok {
		return errors.New(r.Action + " Is not a valid action")
	}

	// make sure there is a valid target DB for the request
	if _, ok := DBS[r.Target]; !ok {
		return errors.New(r.Target + " Is not a valid target")
	}

	return nil
}

// Execute the searchAPI request
func (r *Request) Execute() error {
	HANDLERS[r.Action](r.ResponseWritter, r)
}
