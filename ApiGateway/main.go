package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// middleware merchant
func merchantMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var autor = r.Header.Get("Authoritazion")

		if autor != "merchant" {
			w.Write([]byte("Anda tidak memiliki akses"))
		}

		next.ServeHTTP(w, r)
	}
}

func superAll(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var autor = r.Header.Get("Authoritazion")

		if autor != "su-admin" {
			w.Write([]byte("Anda tidak memiliki akses"))
		}

		next.ServeHTTP(w, r)
	}
}

func getMerchant(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:8000/get-merchant-detail")

	data, _ := ioutil.ReadAll(resp.Body)

	json.NewEncoder(w).Encode(data)
}

func getAllToko(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:8000/get-all-toko")

	data, _ := ioutil.ReadAll(resp.Body)

	json.NewEncoder(w).Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/merchant", merchantMiddle(getMerchant))
	mux.HandleFunc("/toko", merchantMiddle(getAllToko))

	fmt.Println("Server is running")

	http.ListenAndServe(":6000", mux)
}
