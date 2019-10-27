package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func configurePostHeader(t *testing.T, jsonStr []byte) (int, *httptest.ResponseRecorder) {
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

func configureGetHeader(t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getStats)
	handler.ServeHTTP(rr, req)
	return rr
}

func TestMutant0(t *testing.T) {

	var jsonStr = []byte(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)

	status, rr := configurePostHeader(t, jsonStr)

	expected := http.StatusOK

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}

}

func TestMutan1(t *testing.T) {

	var jsonStr = []byte(`{"dna":["ACATACCA","GTGTTACA","GGCATAAG","ATGGGCTC","GTGCCGTA","AAGGGGAG","ATGATGGG","TCCTTCCT"]}`)
	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusOK
	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestMutan2(t *testing.T) {

	var jsonStr = []byte(`{"dna":["TTGTCG","AATAGC","ACGGAA","GCGGAA","CGGGGA","ATCAGA"]}`)
	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusOK
	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestMutan3(t *testing.T) { // Repite la de arriba

	var jsonStr = []byte(`{"dna":["TTGTCG","AATAGC","ACGGAA","GCGGAA","CGGGGA","ATCAGA"]}`)
	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusOK
	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestMutan4(t *testing.T) {

	var jsonStr = []byte(`{"dna":["GCGTCTGCAC","GCGATAGAGG","CGCAATGTTC","CAAGTAAATT","ACTTAAAAGT","TCCCCGGTCC","CGCCTGCACC","GCATTAGTTT","TTCCTACTAG","AATGGACACT"]}`)
	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusOK
	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestMutan5(t *testing.T) {

	var jsonStr = []byte(`{"dna":["AAAACACGCA","ACTGCGAGGA","TGCTGTATTA","TATGCAAGTG","GGGGCAGTCT","CCTAGACGAG","ACAGAACACG","GGATGGGGGT","TTCAACGCTA","TTAGACGATA"]}`)
	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusOK
	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}

func TestHuman0(t *testing.T) {

	var jsonStr = []byte(`{"dna":["CATTGACC","AGAGTGAC","TTAAAGAA","GATTCGGC","TGCTTCGA","CTAAGGCG","TAGCAAAT","ACTAGAAT"]}`)

	status, rr := configurePostHeader(t, jsonStr)

	expected := http.StatusForbidden
	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestHuman1(t *testing.T) {

	var jsonStr = []byte(`{"dna":["TTTA","CTAG","CCAC","AGTC"]}`)

	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusForbidden

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}

func TestHuman2(t *testing.T) {

	var jsonStr = []byte(`{"dna":["TAAGTT","CTCCAA","CCGGTG","GGCCCT","GGCCGC","GGCGCA"]}`)

	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusForbidden

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestHuman3(t *testing.T) {

	var jsonStr = []byte(`{"dna":["GGTCTTGGTA","CGCAATGCAC","GTCAGATGAA","ACGTTACGCG","AGTGCTCACG","CATACCTGCC","TCCGTAAACT","TACGTCTGAC","CAGAACTGTC","TCTTCCAATT"]}`)

	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusForbidden

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}

func TestHuman4(t *testing.T) { // Repite arriba

	var jsonStr = []byte(`{"dna":["GGTCTTGGTA","CGCAATGCAC","GTCAGATGAA","ACGTTACGCG","AGTGCTCACG","CATACCTGCC","TCCGTAAACT","TACGTCTGAC","CAGAACTGTC","TCTTCCAATT"]}`)

	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusForbidden

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestHuman5(t *testing.T) {

	var jsonStr = []byte(`{"dna":["GTGCGGCC","CAATCCCG","GGGATTGT","CCGCTTAA","CTCGGAAC","AGTTTGCG","AGCCAATG","TTACGACT"]}`)

	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusForbidden

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestHuman6(t *testing.T) {

	var jsonStr = []byte(`{"dna":["TCGAGCGTT","GTCATAGGG","CTTATATTT","GAAACTCAT","CCTACACAG","ACGCTAGGG","AGTCTCATC","TATTGAGAC","GGTTAAATG"]}`)

	status, rr := configurePostHeader(t, jsonStr)
	expected := http.StatusForbidden

	if status != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), status)
	}
}
func TestHumanNoValid(t *testing.T) {

	var jsonStr = []byte(`{"dna":["TCGAGCGTT","GTCATAGGG","CTTATATTT","GAAACTCAT","CCTACACA","ACGCTAGGG","AGTCTCATC","TATTGAGAC","GGTTAAATG"]}`)

	status, rr := configurePostHeader(t, jsonStr)
	expected := `{"error":"Secuencia no válida."}`

	if rr.Body.String() != expected || status != 500 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestStats(t *testing.T) {

	rr := configureGetHeader(t)

	expected := `{"count_mutant_dna":5,"count_human_dna":6,"count_ratio_dna":0.8333}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
