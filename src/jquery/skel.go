package main

import (
    "fmt"
    "net/http"
)

func callmeback(w http.ResponseWriter, r *http.Request) {
    hostname := r.URL.Query().Get("error")
    fmt.Fprintf(w, "Hi there, I am 'callmeback'.", hostname)
}

func handler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Fprintf(w, "Hi there, I got a code: %s!", code)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/callmeback", callmeback)
    http.ListenAndServe(":8888", nil)
}
