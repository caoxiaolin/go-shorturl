// Copyright 2016 caoxiaolin

// 一个短链接服务.
package main

import (
	"github.com/caoxiaolin/go-shorturl/utils"
	"database/sql"
    _ "github.com/lib/pq"
	"fmt"
	"log"
	"net/http"
)

const service = "127.0.0.1:4000"

const (
    dbhost      = "127.0.0.1"
    dbport      = 5432
    username    = "rdtest"
    password    = "123456"
    dbname      = "shorturl"
)

type Shorturl struct{}

var db *sql.DB

func init() {
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbhost, dbport, username, password, dbname)
    db, _ = sql.Open("postgres", dsn)
	db.Ping()
}

func (this Shorturl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		uri := r.RequestURI
		l := len(uri)
		res := utils.GetOriUrl(db, uri[1:l])
        log.Printf("[GET] [%s] [%s]", uri, res)
	} else if method == "POST" {
		r.ParseForm()
		res := utils.GetShortUrl(db, r.Form["url"][0])
        log.Printf("[POST] [%s] [%s]", r.Form["url"][0], res)
	}
}

func main() {
	var s Shorturl
	err := http.ListenAndServe(service, s)
	if err != nil {
		log.Fatal(err)
	}
}
