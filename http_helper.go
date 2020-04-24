package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func FixURL(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "http://" + url
	}
	return url
}

func CallHTTP(url string) []byte {
	// Request with get method
	resp, err := http.Get(url)
	if err != nil {
		panic("Couldn't get response with error: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Couldn't get response with error: " + err.Error())
	}
	return body
}
