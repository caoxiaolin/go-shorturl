package serv

import (
	"testing"
)

var (
	turl = "http://www.test.com"
	surl string
)

func TestGetShortUrl(t *testing.T) {
	surl, _ = GetShortUrl(turl)
	if surl == "" {
		t.Error("getShortUrl return empty")
	}
}

func TestGetOriUrl(t *testing.T) {
	if surl != "" {
		ourl, _ := GetOriUrl(surl)
		if ourl != turl {
			t.Error("Expected", surl, ", but got", ourl)
		}
	}
}
