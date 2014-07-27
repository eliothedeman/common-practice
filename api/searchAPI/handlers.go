package searchAPI

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	JSON_FORMATTING_ERROR = "JSON Formatting Error: "
)

// APIHandler interface describes a custom handler for a pre-parsed request
type APIHandler func(http.ResponseWriter, *Request)

// Route is the central routing function for the root endpoint
func Route(w http.ResponseWriter, r *http.Request) {
	req := Request{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(b, &req)
	if err != nil {
		http.Error(w, JSON_FORMATTING_ERROR+err.Error(), 400)
		return
	}
	// Route Request by action type
	HANDLERS[req.Action](w, &req)

}

// Insert is the APIHandler for the insert action
func Insert(w http.ResponseWriter, r *Request) {
	if _, ok := DBS[r.Target]; ok {
		DBS[r.Target].Insert(r.Location, string(r.Data))
	} else {
		http.Error(w, "Target does not exist", 500)
	}

}

// Select is the APIHandler for th select action
func Select(w http.ResponseWriter, r *Request) {
	if _, ok := DBS[r.Target]; ok {
		DBS[r.Target].Select(string(r.Data))
	} else {
		http.Error(w, "target does not exist", 500)
	}

}
