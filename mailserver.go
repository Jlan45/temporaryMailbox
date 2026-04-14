package main

import (
	"fmt"
	"github.com/alash3al/go-smtpsrv"
	"strings"
)

type mailContent struct {
	from    string
	to      string
	title   string
	content string
}

func isAllowedDomain(domain string) bool {
	for _, d := range domains {
		if d == domain {
			return true
		}
	}
	return false
}

func handler(c *smtpsrv.Context) error {
	UserMail := c.To().String()
	UserMail = strings.Trim(UserMail, "<>")
	st := strings.Split(UserMail, "@")
	if len(st) != 2 {
		return fmt.Errorf("Invalid email address")
	}
	s := st[0]
	if !isAllowedDomain(st[1]) {
		return fmt.Errorf("Invalid domain")
	}
	msg, _ := c.Parse()
	content := mailContent{
		from:    strings.Trim(c.From().String(), "<>"),
		title:   msg.Subject,
		content: msg.TextBody,
	}
	if mailBox[s] == nil {
		mailBox[s] = make([]mailContent, 0)
	}
	mailBox[s] = append(mailBox[s], content)
	return nil
}

func startSMTPServer(domainList []string) {
	bannerDomain := "localhost"
	if len(domainList) > 0 {
		bannerDomain = domainList[0]
	}
	cfg := smtpsrv.ServerConfig{
		BannerDomain:    bannerDomain,
		ListenAddr:      ":25",
		MaxMessageBytes: 5 * 1024,
		Handler:         handler,
	}
	fmt.Println(smtpsrv.ListenAndServe(&cfg))
}
