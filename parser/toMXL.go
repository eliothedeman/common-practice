package parser

import (
	"encoding/json"
	"github.com/eliothedeman/go-mxl"
	"log"
)

// ToMXL is an action which takes a []byte and and converts it to an mxl representation
type ToMXL struct {
	Data []byte
}

// Do parses am mxl doc
func (t *ToMXL) Do(m *Manager) {
	doc := &mxl.MXLDoc{}
	err := json.Unmarshal(t.Data, doc)
	if err != nil {
		log.Println(err)
		return
	}
	c := toCommon{Doc: doc}
	m.AddAction(c)
}
