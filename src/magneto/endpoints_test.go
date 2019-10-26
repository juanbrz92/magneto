package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateEntry(t *testing.T) {

	var jsonStr = []byte(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)

	req, err := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checkMutant)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	expected := http.StatusOK
	// expected := http.StatusForbidden

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
