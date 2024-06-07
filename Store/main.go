package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getDetailToko(w http.ResponseWriter, r *http.Request) {
	// get data from db, detail toko

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"nama":           "Altamis Atmaja",
		"jumlah_product": "10",
	})
}

func getAllToko(w http.ResponseWriter, r *http.Request) {
	// get data from db, all toko

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]map[string]string{{"nama": "Altamis Atmaja Toko",
		"jumlah_product": "100"}, {"nama": "Altamis Atmaja Toko 2",
		"jumlah_product": "100"}})
}

func main() {
	var mux = http.NewServeMux()

	mux.HandleFunc("/get-detail-toko", getDetailToko)
	mux.HandleFunc("/get-all-toko", getAllToko)

	fmt.Println("Server is running")

	http.ListenAndServe(":9000", mux)
}
