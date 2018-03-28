package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "%s\n", r.URL.RawQuery)
	q := r.URL.RawQuery
	m, err := url.ParseQuery(q)
	if err != nil {
		log.Fatal(err)
	}
	switch {
	case len(m["name"]) > 0:
		fmt.Fprintf(w, "Hello %s\n", m["name"][0])
	default:
		fmt.Fprintf(w, "%s\n", "[ERROR] please input your name in the request argument, e.g. /?name=pahud")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
