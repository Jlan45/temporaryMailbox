package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func startHTTPServer(Domain string) {
	httpsrv := gin.New()
	httpsrv.GET("/getAddress", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"address": RandStringRunes(8) + "@" + Domain,
		})
	})
	httpsrv.GET("/getMailList/:randomString", func(c *gin.Context) {
		mailHead := c.Param("randomString")
		if mailBox[mailHead] == nil {
			c.JSON(200, gin.H{
				"mails": make([]string, 0),
			})
		} else {
			mails := make([]gin.H, len(mailBox[mailHead]))
			for i, v := range mailBox[mailHead] {
				mails[i] = gin.H{"from": v.from, "title": v.title}
			}
			c.JSON(200, gin.H{
				"mails": mails,
			})
		}
	})
	httpsrv.GET("/getMail/:randomString", func(c *gin.Context) {
		mailHead := c.Param("randomString")
		if mailBox[mailHead] == nil {
			c.JSON(200, gin.H{
				"mail": "没有邮件",
			})
		} else if len(mailBox[mailHead]) == 0 {
			c.JSON(200, gin.H{
				"mail": "没有邮件",
			})
		} else {
			tmpMail := mailBox[mailHead][len(mailBox[mailHead])-1]
			mailBox[mailHead] = mailBox[mailHead][0 : len(mailBox[mailHead])-1]
			c.JSON(200, gin.H{
				"mail": gin.H{
					"from":    tmpMail.from,
					"title":   tmpMail.title,
					"content": tmpMail.content,
				},
			})
		}
	})
	httpsrv.Run(":8090")
}
