package searchAPI

// Request represents an api request

type Request struct {
	Action   string  `json:"action"`
	Data     string  `json:"data"`
	Version  float32 `json:"version"`
	Target   string  `json:"target"`
	Location string  `json:"location"`
}
