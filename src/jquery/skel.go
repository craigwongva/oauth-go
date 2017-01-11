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

func callmebacktomorrow(w http.ResponseWriter, r *http.Request) {
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
fmt.Println("Rain in Williamsburg tomorrow")
fmt.Println(m)
fmt.Println(m["access_token"])
fmt.Println("Snow in Williamsburg tomorrow")
//https://api.github.com/user?access_token=6c7de659eb78b3d5f539559b7f7d40cd5395b025

    resp1, err1 := http.Get("https://api.github.com/user?access_token=6c7de659eb78b3d5f539559b7f7d40cd5395b025")
    if err1 != nil {
        fmt.Println("Something went wrong3 ")
        return
    }

    body1, err1 := 
      ioutil.ReadAll(resp1.Body)
    fmt.Fprintf(w, string(body1))
}

func handler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    fmt.Fprintf(w, "Hi there, I got a code: %s!", code)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/callmebacktomorrow", callmebacktomorrow)
    http.ListenAndServe(":8888", nil)
}
