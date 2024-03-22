package main

import (
	"encoding/json"
	"net/http"
)

func handleRootGet(w http.ResponseWriter, r *http.Request) {
	product := GetAllProduct()
	data, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	}
}

func handleRootPost(w http.ResponseWriter, r *http.Request) {

}

func handleProductGet(w http.ResponseWriter, r *http.Request) {

}

func handleProductPOST(w http.ResponseWriter, r *http.Request) {

}
