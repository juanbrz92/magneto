package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func configureHeader(t *testing.T, jsonStr []byte) (int, *httptest.ResponseRecorder) {
	req, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checkMutant)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	return status, rr
}

func TestMutant0(t *testing.T) {

	var jsonStr = []byte(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)

	status, rr := configureHeader(t, jsonStr)

	expected := http.StatusOK

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}

}
func TestMutan1(t *testing.T) {

	var jsonStr = []byte(`{"dna":["CATTGACC","AGAGTGAC","TTAAAGAA","GATTCGGC","TGCTTCGA","CTAAGGCG","TAGCAAAT","ACTAGAAT"]}`)

	status, rr := configureHeader(t, jsonStr)

	expected := http.StatusForbidden
	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}

func TestMutan2(t *testing.T) {

	var jsonStr = []byte(`{"dna":["ACATACCA","GTGTTACA","GGCATAAG","ATGGGCTC","GTGCCGTA","AAGGGGAG","ATGATGGG","TCCTTCCT"]}`)
	status, rr := configureHeader(t, jsonStr)
	expected := http.StatusOK
	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestHuman1(t *testing.T) {

	var jsonStr = []byte(`{"dna":["TTTA","CTAG","CCAC","AGTC"]}`)

	status, rr := configureHeader(t, jsonStr)
	expected := http.StatusForbidden

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
