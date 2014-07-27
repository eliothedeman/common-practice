package searchAPI

// Response represents an api response

type Response struct {
	Action  string  `json:"action"`
	Data    string  `json:"data"`
	Type    string  `json:"type"`
	Target  string  `json:"target"`
	Version float32 `json:"version"`
}
