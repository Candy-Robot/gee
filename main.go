package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main(){


	//http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//	obj := map[string]interface{}{
	//		"name": "tyc",
	//		"password": "1234",
	//	}
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusOK)
	//	encoder := json.NewEncoder(w)
	//	if err:= encoder.Encode(obj); err != nil {
	//		http.Error(w, err.Error(), 500 )
	//	}
	//})
	//http.ListenAndServe(":9999", nil)
	eng := gee.New()
	eng.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL = %q\n", req.URL.Path)
	})
	eng.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header{
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	eng.RUN(":9999")
}