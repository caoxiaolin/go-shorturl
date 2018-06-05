// Copyright 2016 caoxiaolin

// 一个短链接服务.
package main

import (
	"database/sql"
	"fmt"
	"github.com/caoxiaolin/go-shorturl/config"
	"github.com/caoxiaolin/go-shorturl/utils"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Shorturl struct{}

var (
	address string
	db      *sql.DB
	cfg     *config.TomlConfig
)

func init() {
	cfg = config.Load()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.UserName, cfg.Database.PassWord, cfg.Database.DbName)
	db, _ = sql.Open("postgres", dsn)
	db.SetMaxOpenConns(cfg.Database.MaxConn)
	db.Ping()
}

func (this Shorturl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		uri := r.RequestURI
		l := len(uri)
		res := utils.GetOriUrl(db, uri[1:l])
		if res != "" {
			fmt.Fprintln(w, res)
			log.Printf("[GET] [%s] [%s] [%s]", r.RemoteAddr, uri, res)
		} else {
			w.WriteHeader(404)
			log.Printf("[GET] [%s] [%s] [%s]", r.RemoteAddr, uri, "404 NOT FOUND")
		}
	} else if method == "POST" {
		r.ParseForm()
		res := utils.GetShortUrl(db, r.Form["url"][0])
		fmt.Fprintln(w, "http://"+address+"/"+res)
		log.Printf("[POST] [%s] [%s] [%s]", r.RemoteAddr, r.Form["url"][0], res)
	}
}

func main() {
	var s Shorturl
	address = fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("service starting on " + address + " ...")
	err := http.ListenAndServe(address, s)
	if err != nil {
		log.Fatal(err)
	}
}
