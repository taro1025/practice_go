// +build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handlerFafa(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, " %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handlerFafa)
  //http.HandleFunc("/", handler) //tell the http package to handle all requests to the webroot with handler
  log.Fatal(http.ListenAndServe(":8080", nil))
}