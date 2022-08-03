package main

import (
	"fmt"
	URL "net/url"
	"testing"
)

func TestURL(t *testing.T) {
	info, err := URL.Parse("https://a.com")
	fmt.Println(info, err)
	fmt.Println(info.Scheme, err)
	url := "http://a.com/path<a>/to?query%3C=<script>&b=<#hash=<div>"
	fmt.Println(encodeURI(url))

}

func encodeURI(url string) string {
	u, _ := URL.Parse(url)
	u.RawQuery = u.Query().Encode()
	return u.String()
}
