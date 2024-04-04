package main

import "net/http"

func setNoChacheHeaders(req *http.Request) {
	req.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Expires", "0")
}
