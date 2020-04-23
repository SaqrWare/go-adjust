package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strings"
)

func FixURLs(params []string) []string {
	result := make([]string, 0)
	for _, item := range params {
		if !strings.HasPrefix(item, "http://") && !strings.HasPrefix(item, "https://") {
			item = "http://" + item
		}
		result = append(result, item)
	}
	return result
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

func HashResponse(url string, body []byte) (string, string) {
	hasher := md5.New()
	hasher.Write(body)
	return url, hex.EncodeToString(hasher.Sum(nil))
}
