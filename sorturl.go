package main

import (
    "fmt"
    "log"
    "net/http"
    "./utils/"
)

type Sorturl struct{}

func (this Sorturl) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    method := r.Method
    if method == "GET" {
        uri := r.RequestURI
        l := len(uri)
        fmt.Fprint(w, utils.GetOriUrl(uri[1:l]))
        fmt.Fprint(w, "\n")
    } else if method == "POST" {
        r.ParseForm()
        fmt.Fprint(w, utils.GetSortUrl(r.Form["url"][0]))
        fmt.Fprint(w, "\n")
    }
}

func main() {
    var s Sorturl
    err := http.ListenAndServe("192.168.245.128:4000", s)
    if err != nil {
        log.Fatal(err)
    }
}
