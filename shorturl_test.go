package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var uri string

/**
 * 测试获取一个不存在的短链，应该返回404
 */
func TestGetNonexistentUrl(t *testing.T) {
	req, err := http.NewRequest("GET", "/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ShorturlServer(rr, req)

	if rr.Code != 404 {
		t.Error("Http code expected 404, but got ", rr.Code)
	}
}

/**
 * 测试生成一个新的短链
 */
func TestSeturl(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader("url=http://www.google.com"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	ShorturlServer(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	u := strings.TrimSpace(string(body))
	if u == "" {
		t.Error("Expected http://"+address, "/..., but got ", u)
	}

	urlParse, _ := url.Parse(u)
	uri = urlParse.Path
}

/**
 * 测试获取已经存在的短链
 */
func TestGetExistedUrl(t *testing.T) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Cookie", "debug=1")
	rr := httptest.NewRecorder()
	ShorturlServer(rr, req)

	if rr.Code != 200 {
		t.Error("Http code expected 200, but got ", rr.Code)
	}

	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	u := strings.TrimSpace(string(body))
	if u != "http://www.google.com" {
		t.Error("Expected http://www.google.com, but got", u)
	}
}

/**
 * 测试获取短链并跳转
 */
func TestRedirectUrl(t *testing.T) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ShorturlServer(rr, req)

	if rr.Code != 302 {
		t.Error("Http code expected 302, but got ", rr.Code)
	}
}
