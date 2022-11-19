package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	targetUrl = "https://proxy.com"
	Port      = "8080"
)

func main() {
	// start server
	http.HandleFunc("/", proxyPass)
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}

func proxyPass(res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(targetUrl)
	proxy := httputil.NewSingleHostReverseProxy(url)
	req.Header.Set("X-Some-Header", "TBD")
	proxy.ServeHTTP(res, req)
}
