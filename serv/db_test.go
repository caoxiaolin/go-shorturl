package serv

import (
	"testing"
)

var (
	turl = "http://www.test.com"
	surl string
)

func TestGetShortUrl(t *testing.T) {
	surl, err := GetShortUrl("")
	if err == nil {
		t.Error("Expected empty, but got", surl)
	}

	surl, _ = GetShortUrl(turl)
	if surl == "" {
		t.Error("getShortUrl return empty")
	}
}

func TestGetOriUrl(t *testing.T) {
	ourl, _ := GetOriUrl("0")
	if ourl != "" {
		t.Error("Expected empty, but got", ourl)
	}

	if surl != "" {
		ourl, _ = GetOriUrl(surl)
		if ourl != turl {
			t.Error("Expected", surl, ", but got", ourl)
		}
	}
}
