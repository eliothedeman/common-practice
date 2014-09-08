package api

// Request is the generic api handler interface
type Request interface {
	Execute() error
	Validate() error
	Respond() error
}
