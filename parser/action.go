package parser

// Action is the basic interface for tasks the manager handles
type Action interface {
	Do(*Manager)
}
