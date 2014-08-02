package parser

// Manager manages jobs as they move through the parsing pipeline
type Manager struct {
	workChan chan Action
}

// NewManager creates and returns a new manager
func NewManager() *Manager {
	m := &Manager{}
	m.workChan = make(chan Action, 1000)
	go m.Manage()
	return m
}

// AddAction addes a new action to the workChan to be picked up by the manager
func (m *Manager) AddAction(a *Action) {
	m.workChan <- a
}

// Manage is the main loop the manager uses to do work
func (m *Manager) Manage() error {
	for {
		select {
		case a := <-m.workChan:
			go a.Do()
		}
	}
}
