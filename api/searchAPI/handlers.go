package searchAPI

import (
	"encoding/json"
	"errors"
)

// APIHandler interface describes a custom handler for a pre-parsed request
type APIHandler func(*Request) (map[string]interface{}, error)

// Insert is the APIHandler for the insert action
func Insert(r *Request) (map[string]interface{}, error) {
	if _, ok := DBS[r.Target]; ok {
		err := DBS[r.Target].CreateAndUpdateIndicies(r.Data, r.Location)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Target " + r.Target + " Not found")
	}
	return map[string]interface{}{
		"version":   r.Version,
		"localtion": r.Location,
		"status":    "success",
	}, nil

}

// Select is the APIHandler for th select action
func Select(r *Request) (map[string]interface{}, error) {
	if _, ok := DBS[r.Target]; ok {
		i := DBS[r.Target].Select(r.Data)
		// if the returned value is nil (doesn't exist).
		// return an error
		if i == nil {
			return nil, errors.New(r.Data + " Not found in search database")
		}
		b, err := json.Marshal(i)
		if err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"version":  r.Version,
			"location": r.Location,
			"status":   "success",
			"data":     b,
		}, nil

	}
	return nil, errors.New("Target " + r.Target + " does not exist")

}
