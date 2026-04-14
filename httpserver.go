package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
var htmlIndex = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>随机邮箱生成器</title>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #f0f0f0;
            margin: 0;
            padding: 0;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            height: 100vh;
        }

        h1 {
            color: #333;
        }

        .buttons {
            margin-top: 20px;
        }

        button {
            padding: 10px 20px;
            margin: 0 10px;
            font-size: 16px;
            cursor: pointer;
            background-color: #4CAF50;
            color: white;
            border: none;
            border-radius: 5px;
        }

        button:hover {
            background-color: #45a049;
        }

        select {
            padding: 10px;
            font-size: 16px;
            border: 1px solid #ddd;
            border-radius: 5px;
            margin-top: 20px;
        }

        #randomAddress {
            margin-top: 20px;
            padding: 10px;
            border: 1px solid #ddd;
            border-radius: 5px;
            background-color: #fff;
        }

        #result {
            margin-top: 20px;
            padding: 10px;
            border: 1px solid #ddd;
            border-radius: 5px;
            background-color: #fff;
        }
    </style>
</head>
<body>
    <h1>随机邮箱生成器</h1>
    <div>
        <label for="domainSelect">选择域名: </label>
        <select id="domainSelect"></select>
    </div>
    <div class="buttons">
        <button id="getRandom" >获取随机邮件地址</button>
        <button id="getMailList" >获取邮件列表</button>
        <button id="getMail" >获取邮件</button>
      </div>
    </div>
        <div id="randomAddress"></div>
    <div id="result"></div>
    <script>
        let randomMailAddress = ''; // 保存随机邮件地址

        // 页面加载时获取可用域名列表
        fetch('/getDomains')
            .then(response => response.json())
            .then(data => {
                const select = document.getElementById('domainSelect');
                data.domains.forEach(domain => {
                    const option = document.createElement('option');
                    option.value = domain;
                    option.textContent = domain;
                    select.appendChild(option);
                });
            })
            .catch(error => console.error('Error:', error));

        // 获取随机邮件地址
        document.getElementById('getRandom').addEventListener('click', function() {
            const domain = document.getElementById('domainSelect').value;
            if (!domain) {
                alert('请先选择一个域名！');
                return;
            }
            // 调用后端接口，这里假设使用 fetch
            fetch(` + "`/getAddress?domain=${encodeURIComponent(domain)}`" + `)
                .then(response => response.json())
                .then(data => {
                    randomMailAddress = data.random; // 保存随机邮件地址
                    document.getElementById('randomAddress').innerText = '随机邮件地址: ' + data.address;
                })
                .catch(error => console.error('Error:', error));
        });

        // 获取邮件列表
        document.getElementById('getMailList').addEventListener('click', function() {
            if (!randomMailAddress) {
                alert('请先获取随机邮件地址！');
                return;
            }

            // 调用后端接口，这里假设使用 fetch
            fetch(` + "`/getMailList/${randomMailAddress}`" + `)
                .then(response => response.json())
                .then(data => {
                    // 假设展示第一封邮件的标题
                    document.getElementById('result').innerHTML = '';
                    if(data.mails.length==0){
                        document.getElementById('result').innerText = '邮件还没收到，耐心等一下哦';
                    }
else{
		data.mails.forEach(mail => {
                        const mailElement = document.createElement('div');
                        mailElement.innerText = ` + "`发件人: ${mail.from}, 标题: ${mail.title}`" + `;
                        document.getElementById('result').appendChild(mailElement);
                    });}

                })
                .catch(error => console.error('Error:', error));
        });

        // 获取邮件内容
        document.getElementById('getMail').addEventListener('click', function() {
            if (!randomMailAddress) {
                alert('请先获取随机邮件地址！');
                return;
            }

            // 调用后端接口，这里假设使用 fetch
            fetch(` + "`/getMail/${randomMailAddress}`" + `)
                .then(response => response.json())
                .then(data => {
                    // 假设展示邮件内容
                    document.getElementById('result').innerText = '邮件内容: ' + data.mail.content;
                })
                .catch(error => console.error('Error:', error));
        });
    </script>
</body>
</html>`

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func startHTTPServer(domainList []string) {
	httpsrv := gin.New()
	httpsrv.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, htmlIndex)
	})
	httpsrv.GET("/getDomains", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"domains": domainList,
		})
	})
	httpsrv.GET("/getAddress", func(c *gin.Context) {
		domain := c.Query("domain")
		if domain == "" && len(domainList) > 0 {
			domain = domainList[0]
		}
		if !isAllowedDomain(domain) {
			c.JSON(400, gin.H{
				"error": "Invalid domain",
			})
			return
		}
		tmp := RandStringRunes(8)
		c.JSON(200, gin.H{
			"random":  tmp,
			"address": tmp + "@" + domain,
			"domain":  "@" + domain,
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
	httpsrv.Run(":80")
}
