package searchAPI

import (
	"encoding/json"
	"github.com/eliothedeman/common-practice/error_codes"
	"io"
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
	req := &Request{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, JSON_FORMATTING_ERROR+err.Error(), error_codes.API_INVALID_REQUEST)
		return
	}
	err = json.Unmarshal(b, &req)
	if err != nil {
		http.Error(w, JSON_FORMATTING_ERROR+err.Error(), error_codes.API_INVALID_REQUEST)
		return
	}
	// validate the request
	err = req.validate()
	if err != nil {
		http.Error(w, JSON_FORMATTING_ERROR+err.Error(), error_codes.API_INVALID_REQUEST)
		return
	}
	// Route Request by action type
	HANDLERS[req.Action](w, req)

}

// Insert is the APIHandler for the insert action
func Insert(w http.ResponseWriter, r *Request) {
	if _, ok := DBS[r.Target]; ok {
		err := DBS[r.Target].CreateAndUpdateIndicies(r.Data, r.Location)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		log.Println(DBS[r.Target].Select(r.Data))
		http.Error(w, "Target does not exist", 400)
	}

}

// Select is the APIHandler for th select action
func Select(w http.ResponseWriter, r *Request) {
	if _, ok := DBS[r.Target]; ok {
		i := DBS[r.Target].Select(r.Data)
		// if the returned value is nil (doesn't exist).
		// return an empty json set to the client.
		if i == nil {
			io.WriteString(w, "{}")
			return
		}
		b, err := json.Marshal(i)
		if err != nil {
			http.Error(w, "Bad Data "+err.Error(), 500)
			log.Println(err)
		} else {
			_, err := io.WriteString(w, string(b))
			if err != nil {
				http.Error(w, "", 400)
			} else {
				return
			}
		}

	} else {
		http.Error(w, "target does not exist", 400)
	}

}

// APIError sends back an API Error to the requester
