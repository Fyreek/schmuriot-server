package models

import "encoding/json"

// BaseAction is the base type for all actions
type BaseAction struct {
	Action string `json:"action"`
}

// Marshal makes json out of the BaseAction struct and gives back a byte array
func (ba *BaseAction) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(ba)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Unmarshal fills the struct from a json byte array
func (ba *BaseAction) Unmarshal(bytes []byte) error {
	err := json.Unmarshal(bytes, ba)
	return err
}

// Check checks for the provided action
func (ba *BaseAction) Check(action string) bool {
	if ba.Action == action {
		return true
	}
	return false
}
