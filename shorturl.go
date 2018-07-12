// Copyright 2016 caoxiaolin

// 一个短链接服务.
package main

import (
	"fmt"
	"github.com/caoxiaolin/go-shorturl/config"
	"github.com/caoxiaolin/go-shorturl/utils"
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
func handlePost(w http.ResponseWriter, r *http.Request){
        postUrl := utils.GetPostUrl(r)
		res, err := utils.GetShortUrl(postUrl)
        if err != nil {
            res = err.Error()
            w.WriteHeader(http.StatusBadRequest)
		    fmt.Fprintln(w, res)
        } else {
		    fmt.Fprintln(w, "http://"+address+"/"+res)
        }
		utils.Logger.Printf("[POST] [%s] [%s] [%s]", r.RemoteAddr, postUrl, res)

}

// handleGet can handle get request
func handleGet(w http.ResponseWriter, r *http.Request){
		uri := r.URL.Path
		l := len(uri)
		res := utils.GetOriUrl(uri[1:l])
		if res != "" {
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
