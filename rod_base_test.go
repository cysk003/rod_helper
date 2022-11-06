package rod_helper

import (
	"testing"
	"time"
)

func TestNewBrowserBase(t *testing.T) {

	httpProxyUrl := "http://192.168.50.252:20171"
	movieUrl := "https://www.google.com"
	b, err := NewBrowserBase("", httpProxyUrl, true, true)
	if err != nil {
		t.Fatal(err)
	}
	page, _, err := NewPageNavigateWithProxy(b, httpProxyUrl, movieUrl, 15*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	println(page.MustHTML())
}
