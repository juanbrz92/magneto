package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"middlewares/basicauthmiddleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var host = "http://localhost"
var port = "12345"
var connectionString = "juanbrz92:36554061jjhjjh@tcp(127.0.0.1:3306)/magneto?charset=utf8&parseTime=True&loc=Local"

func main() {
	var router *mux.Router
	router = mux.NewRouter().StrictSlash(true)
	router.Handle("/api/mutant", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(checkMutant))).Methods("POST")
	router.Handle("/api/stats", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(getStats))).Methods("GET")
	fmt.Println("Listening on port :12345")
	http.ListenAndServe(":"+port, router)
}

func checkMutant(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var sequence sequence
	err = decoder.Decode(&sequence)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}
	statement, err := db.Prepare("insert into sequence (dna, result) values(?,?)")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}
	defer statement.Close()

	resultTest := isMutant(sequence.DNA[:])

	data, err := json.Marshal(sequence.DNA)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}
	res, err := statement.Exec(data, resultTest)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "There was problem entering the sequence.")
		return
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		id, _ := res.LastInsertId()
		sequence.ID = int(id)
		if resultTest {
			sequence.RESULT = "Es mutante"
			respondWithJSON(w, http.StatusOK, sequence)
		} else {
			sequence.RESULT = "No es mutante"
			respondWithJSON(w, http.StatusForbidden, sequence)
		}
	}
}

func getStats(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}
	var countMutantDna sql.NullInt32
	var countHumanDna sql.NullInt32
	var ratio sql.NullFloat64
	query := `
		SELECT 
		COUNT(CASE WHEN s.result = '1' THEN 1 ELSE NULL END) AS countMutantDna,
		COUNT(CASE WHEN s.result = '0' THEN 1 ELSE NULL END) AS countHumanDna,
		COUNT(CASE WHEN s.result = '1' THEN 1 ELSE NULL END)/COUNT(CASE WHEN s.result = '0' THEN 1 ELSE NULL END) AS ratio
		FROM sequence s`
	err = db.QueryRow(query).Scan(&countMutantDna, &countHumanDna, &ratio)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong.")
		return
	}
	var stats stat
	stats.CountMutantDna = countMutantDna.Int32
	stats.CountHumanDna = countHumanDna.Int32
	stats.Ratio = ratio.Float64
	respondWithJSON(w, http.StatusOK, stats)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
