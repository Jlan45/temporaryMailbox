package main

import "os"

var mailBox = make(map[string][]mailContent)
var subDomain = os.Getenv("SUBDOMAIN")

func main() {
	go startHTTPServer(subDomain)
	startSMTPServer(subDomain)
}
