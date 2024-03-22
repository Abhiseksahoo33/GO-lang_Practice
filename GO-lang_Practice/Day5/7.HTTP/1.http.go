/*
query path

	mav.com/employee GET, POST, PUT, DELETE

query parameter [http://google.com/search?.....]
Headers- (Contente-type: application/json, content-length:1024, )
body
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Http server")

	//http://localhost:1234
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We have hit / path " + r.Method))
		w.WriteHeader(http.StatusOK)
	})

	//http://localhost:1234/test
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We have hit /test path " + r.Method))
		w.WriteHeader(http.StatusOK)
	})

	//http://localhost:1234/test/case1
	http.HandleFunc("/test/case1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We have hit /test/case1 path " + r.Method))
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":1234", nil)
}
