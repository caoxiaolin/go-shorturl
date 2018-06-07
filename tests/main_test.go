package tests

import (
	"fmt"
	"github.com/caoxiaolin/go-shorturl/config"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var (
	address string
	cfg     *config.TomlConfig
)

func init() {
	cfg = config.Load()
	address = fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
}

func TestGeturl(t *testing.T) {
	resp, err := http.Get("http://" + address + "/1")
	if err != nil {
		t.Error("http server error")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		t.Error("http code expected 404, but got ", resp.StatusCode)
	}
}

func TestSeturl(t *testing.T) {
	resp, err := http.Post("http://"+address,
		"application/x-www-form-urlencoded",
		strings.NewReader("url=http://www.google.com"))
	if err != nil {
		t.Error("http server error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("http server error")
	}

	response := string(body)
	if response != "http://"+address+"/1" {
		t.Error("Expected http://"+address+"/1, but got ", response)
	}
}
