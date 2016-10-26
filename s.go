package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/abc/", func(w http.ResponseWriter, r *http.Request) {
		print(".")
	
		w.Write([]byte("yo!"))
	}) 
	
	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic(err)
	}
}