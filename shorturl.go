// Copyright 2016 caoxiaolin

// 一个短链接服务.
package main

import (
	"fmt"
	"github.com/caoxiaolin/go-shorturl/config"
	"github.com/caoxiaolin/go-shorturl/serv"
	"github.com/caoxiaolin/go-shorturl/utils"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

var (
	address string
)

func init() {
	address = fmt.Sprintf("%s:%d", config.Cfg.Server.Host, config.Cfg.Server.Port)
}

// ShorturlServer handle post or get requests
func ShorturlServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handleGet(w, r)
	} else if r.Method == "POST" {
		handlePost(w, r)
	}
}

// handlePost can handle post request
func handlePost(w http.ResponseWriter, r *http.Request) {
	postUrl := utils.GetPostUrl(r)
	res, err := serv.GetShortUrl(postUrl)
	if err != nil {
		res = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, res)
	} else {
		rdskey := "su_" + res
		serv.Rds.Do("SET", rdskey, postUrl, "EX", 86400)
		fmt.Fprintln(w, "http://"+address+"/"+res)
	}
	utils.Logger.Printf("[POST] [%s] [%s] [%s]", r.RemoteAddr, postUrl, res)

}

// handleGet can handle get request
func handleGet(w http.ResponseWriter, r *http.Request) {
	var res string
	uri := r.URL.Path
	l := len(uri)
	rdskey := "su_" + uri[1:l]
	if rdsval, _ := redis.String(serv.Rds.Do("GET", rdskey)); rdsval != "" {
		res = rdsval
	} else {
		res, _ = serv.GetOriUrl(uri[1:l])
	}
	if res != "" {
		serv.Rds.Do("SET", rdskey, res, "EX", 86400, "NX")
		//debug mode
		debug, _ := r.Cookie("debug")
		if debug != nil && debug.Value == "1" {
			fmt.Fprintln(w, res)
		} else {
			http.Redirect(w, r, res, http.StatusFound)
		}
		utils.Logger.Printf("[GET] [%s] [%s] [%s]", r.RemoteAddr, uri, res)
	} else {
		http.NotFound(w, r)
		utils.Logger.Printf("[GET] [%s] [%s] [404 NOT FOUND]", r.RemoteAddr, uri)
	}

}

func main() {
	log.Printf("Service starting on %s ...", address)
	http.HandleFunc("/", ShorturlServer)
	utils.Logger.Fatal(http.ListenAndServe(address, nil))
}
