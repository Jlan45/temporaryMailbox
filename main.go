package main

import (
	"os"
	"strings"
)

var mailBox = make(map[string][]mailContent)
var domains []string

func init() {
	// Support multiple domains via SUBDOMAINS (comma-separated), with backward compatibility for SUBDOMAIN
	if env := os.Getenv("SUBDOMAINS"); env != "" {
		for _, d := range strings.Split(env, ",") {
			d = strings.TrimSpace(d)
			if d != "" {
				domains = append(domains, d)
			}
		}
	}
	if len(domains) == 0 {
		if env := os.Getenv("SUBDOMAIN"); env != "" {
			domains = append(domains, strings.TrimSpace(env))
		}
	}
}

func main() {
	go startHTTPServer(domains)
	startSMTPServer(domains)
}
