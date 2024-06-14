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
			return
		}

		next.ServeHTTP(w, r)
	}
}

func superMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var autor = r.Header.Get("Authoritazion")

		if autor != "su-admin" {
			w.Write([]byte("Anda tidak memiliki akses"))
			return
		}

		next.ServeHTTP(w, r)
	}
}

type Merchant struct {
	Nama          string `json:"nama`
	NamaToko      string `json:"nama_toko`
	jumlahProduct int    `json:"jumlah_product`
}

type Toko struct {
	NamaToko      string `json:"nama_toko`
	jumlahProduct int    `json:"jumlah_product`
}

func getMerchant(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:8000/api/get-merchant-detail")

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	var merch = &Merchant{}
	json.Unmarshal(data, &merch)

	json.NewEncoder(w).Encode(merch)
}

func getAllToko(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:8000/api/get-all-toko")

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	var toko = &[]Toko{}
	json.Unmarshal(data, &toko)

	json.NewEncoder(w).Encode(toko)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/merchant", merchantMiddle(getMerchant))
	mux.HandleFunc("/toko", superMiddle(getAllToko))

	fmt.Println("Server is running")

	http.ListenAndServe(":6000", mux)
}
