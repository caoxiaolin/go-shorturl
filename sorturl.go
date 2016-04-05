// Copyright 2016 caoxiaolin

// 一个短链接服务.
package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    "./utils/"
)

const host = "192.168.245.128:4000"

type Sorturl struct{}

var db *sql.DB

func init() {
    dsn := "root:123456@tcp(192.168.245.128:3550)/sorturl?charset=utf8"
    db, _ = sql.Open("mysql", dsn)
    db.SetMaxOpenConns(2000)
    db.SetMaxIdleConns(1000)
    db.Ping()
}

func (this Sorturl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    method := r.Method
    if method == "GET" { 
        uri := r.RequestURI
        l := len(uri)
        fmt.Fprint(w, utils.GetOriUrl(db, uri[1:l]))
        fmt.Fprint(w, "\n")
    } else if method == "POST" {
        r.ParseForm()
        fmt.Fprint(w, "http://" + host + "/" + utils.GetSortUrl(db, r.Form["url"][0]))
        fmt.Fprint(w, "\n")
    }
}

func main() {
    var s Sorturl
    err := http.ListenAndServe(host, s)
    if err != nil {
        log.Fatal(err)
    }
}
