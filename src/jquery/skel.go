package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "strings"
)

func callmeback(w http.ResponseWriter, r *http.Request) {
    hostname := r.URL.Query().Get("error")
    fmt.Fprintf(w, "Hi there, I am 'callmeback'.", hostname)
}

func handler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    fmt.Fprintf(w, "Hi there, I got a code: %s!", code)

    resp, err := http.PostForm("https://github.com/login/oauth/access_token",
    url.Values{"code": {code}, "client_id": {os.Args[1]}, "client_secret": {os.Args[2]}})
    if err != nil {
        fmt.Println("Something went wrong2 ")
        return
    }

    body, err := 
      ioutil.ReadAll(resp.Body)
    fmt.Fprintf(w, string(body))
    a := string(body)
    	


var m map[string]string
var ss []string

//s := "A=B&C=D&E=F"
ss = strings.Split(a, "&")
m = make(map[string]string)
for _, pair := range ss {
    z := strings.Split(pair, "=")
    m[z[0]] = z[1]
}
fmt.Println("Rain in Williamsburg")
fmt.Println(m)
fmt.Println("Snow in Williamsburg")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/callmeback", callmeback)
    http.ListenAndServe(":8888", nil)
}
