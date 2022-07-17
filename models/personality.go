package models

type Personalidade struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty" json:"name,omitempty"`
	History string `json:"history,omitempty" json:"history,omitempty"`
}

var Personalidades []Personalidade
