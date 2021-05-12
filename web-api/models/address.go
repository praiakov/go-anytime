package models

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
