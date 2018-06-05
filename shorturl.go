// Copyright 2016 caoxiaolin

// 一个短链接服务.
package main

import (
	"database/sql"
	"fmt"
	"github.com/caoxiaolin/go-shorturl/utils"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const service = "127.0.0.1:4000"

const (
	dbhost   = "127.0.0.1"
	dbport   = 5432
	username = "rdtest"
	password = "123456"
	dbname   = "shorturl"
)

type Shorturl struct{}

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbhost, dbport, username, password, dbname)
	db, _ = sql.Open("postgres", dsn)
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func (this Shorturl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		uri := r.RequestURI
		l := len(uri)
		res := utils.GetOriUrl(db, uri[1:l])
		fmt.Fprintln(w, res)
		log.Printf("[GET] [%s] [%s]", uri, res)
	} else if method == "POST" {
		r.ParseForm()
		res := utils.GetShortUrl(db, r.Form["url"][0])
		fmt.Fprintln(w, "http://"+service+"/"+res)
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
