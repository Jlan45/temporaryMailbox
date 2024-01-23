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

func handler(c *smtpsrv.Context) error {
	UserMail := c.To().String()
	UserMail = strings.Trim(UserMail, "<>")
	st := strings.Split(UserMail, "@")
	s := st[0]
	if st[1] != subDomain {
		return fmt.Errorf("Invalid domain")
	}
	//msg, _ := mail.ReadMessage(c)
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
func startSMTPServer(subDomain string) {
	cfg := smtpsrv.ServerConfig{
		BannerDomain:    subDomain,
		ListenAddr:      ":25",
		MaxMessageBytes: 5 * 1024,
		Handler:         handler,
	}
	fmt.Println(smtpsrv.ListenAndServe(&cfg))
}
