// Copyright 2016 caoxiaolin

// 一个短链接服务.
package main

import (
	"github.com/caoxiaolin/go-shorturl/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

const host = "127.0.0.1:4000"

type Shorturl struct{}

var db *sql.DB

func init() {
	dsn := "root:123456@tcp(192.168.245.128:3550)/shorturl?charset=utf8"
	db, _ = sql.Open("mysql", dsn)
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
		fmt.Println("[" + time.Now().Format("2006/01/02 15:04:05") + "] [GET] [" + uri + "] [" + res + "]")
		fmt.Fprint(w, res)
		fmt.Fprint(w, "\n")
	} else if method == "POST" {
		r.ParseForm()
		res := utils.GetShortUrl(db, r.Form["url"][0])
		fmt.Println("[" + time.Now().Format("2006/01/02 15:04:05") + "] [POST] [" + r.Form["url"][0] + "] [" + res + "]")
		fmt.Fprint(w, "http://"+host+"/"+res)
		fmt.Fprint(w, "\n")
	}
}

func main() {
	var s Shorturl
	err := http.ListenAndServe(host, s)
	if err != nil {
		log.Fatal(err)
	}
}
