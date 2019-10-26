package main

type sequence struct {
	ID     int      `json:"id,omitempty"`
	DNA    []string `json:"dna,omitempty"`
	RESULT string   `json:"result,omitempty"`
}
