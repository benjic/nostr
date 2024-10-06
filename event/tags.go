package event

import (
	"encoding/json"
)

type Tag []string

func (p Tag) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]string(p))
}

type Tags []Tag

func (p Tags) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Tag(p))
}
