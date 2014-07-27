package api

// Request represents an api request

type Request struct {
	Action   string  `json:"action"`
	Data     []uint8 `json:"data"`
	Version  float32 `json:"version"`
	Target   string  `json:"target"`
	Location string  `json:"location"`
}
