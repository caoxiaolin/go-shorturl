package main

import (
	"fmt"
	"github.com/caoxiaolin/go-shorturl/config"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var url string

func init() {
	cfg = config.Load("./config/")
	address = fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
}

func TestGetNonexistentUrl(t *testing.T) {
	resp, err := http.Get("http://" + address + "/0")
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		t.Error("Http code expected 404, but got ", resp.StatusCode)
	}
}

func TestSeturl(t *testing.T) {
	resp, err := http.Post("http://"+address, "application/x-www-form-urlencoded", strings.NewReader("url=http://www.google.com"))
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	url = strings.TrimSpace(string(body))
	if url == "" {
		t.Error("Expected http://", address, "/..., but got empty")
	}
}

func TestGetExistedUrl(t *testing.T) {
	client := &http.Client{}

	// Declare HTTP Method and Url
	req, err := http.NewRequest("GET", url, nil)

	// Set cookie
	req.Header.Set("Cookie", "debug=1")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	res := strings.TrimSpace(string(body))
	if res != "http://www.google.com" {
		t.Error("Expected http://www.google.com, but got ", res)
	}
}
