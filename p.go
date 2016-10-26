package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	var addr string

	flag.StringVar(&addr, "p", ":1234", "...")
	flag.Parse()

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:1234",
		Path:   "/abc/",
	})

	log.Fatal(http.ListenAndServe(":1235", proxy))
}
