# 简易临时邮箱生成
go buid即可使用，请将自己的域名写入环境变量的SUBDOMAIN中
```
/getRandom
{
    "address": "79lb91nk@xxx.com"
}
/getMailList/79lb91nk
{
    "mails": [
        {
            "from": "from@xxx.com",
            "title": "SMTP 邮件测试"
        }
    ]
}
/getMail/79lb91nk
{
    "mail": {
        "content": "邮件发送测试...",
        "from": "from@xxx.com",
        "title": "SMTP 邮件测试"
    }
}
```
调用getMail后对应的邮件会被pop掉，无需担心大量邮件堆积占满内存<br>
测试使用
http://mail.nothinglikethis.asia/
