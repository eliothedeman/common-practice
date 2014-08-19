package searchAPI

import (
	"errors"
)

// Request represents an api request
type Request struct {
	Action   string  `json:"action"`
	Data     string  `json:"data"`
	Version  float32 `json:"version"`
	Target   string  `json:"target"`
	Location string  `json:"location"`
}

// validate is a sanity check for a request. If request is valid, return nil
func (r *Request) validate() error {

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
