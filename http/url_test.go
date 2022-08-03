package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParse(t *testing.T) {
	pn := fmt.Println
	m, _ := url.ParseQuery(`x=1&y=2&y=3;z`)
	pn(m)

	u, _ := url.Parse("http://bing.com:8080/search?q1=dotnet")
	fmt.Printf("url.Parse:%#v\n", u)
	u.Scheme = "https"
	u.Host = "google.com" // Hostname() + Port()
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u.String())
}

func TestUrlValues(t *testing.T) {
	forms := url.Values{}
	forms.Add("a", "1")
	data := forms.Encode()

	if data != "a=1" {
		t.Fatal("invalid url values")
	}
}
