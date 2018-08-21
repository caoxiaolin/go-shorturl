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

func TestGetNoExistedOriUrl(t *testing.T) {
	ourl, _ := GetOriUrl("0")
	if ourl != "" {
		t.Error("Expected empty, but got", ourl)
	}
}

func TestGetOriUrl(t *testing.T) {
	if surl != "" {
		ourl, err := GetOriUrl(surl)
		if err != nil {
			t.Error(err)
		}
		if ourl != turl {
			t.Error("Expected", surl, ", but got", ourl)
		}
	}
}
