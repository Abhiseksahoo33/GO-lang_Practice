package main

import "github.com/gorilla/mux"

func CraateCatlogMux() *mux.Router {
	mx := mux.NewRouter()
	mx.HandleFunc("/", handleRootGet).Methods("GET")
	mx.HandleFunc("/", handleRootPost).Methods("POST")
	mx.HandleFunc("/product", handleProductGet).Methods("GET")
	mx.HandleFunc("/product", handleProductPOST).Methods("POST")
	return mx
}
