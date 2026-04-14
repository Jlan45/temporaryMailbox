# 简易临时邮箱

一个用 Go 编写的轻量级临时邮箱服务，内置 SMTP 服务器（端口 25）和 HTTP 接口（端口 80），可自托管在你自己的域名下。

## 前置要求

- Go 1.18+
- 一个你拥有解析权限的域名
- 服务器开放 **25**（SMTP）和 **80**（HTTP）端口
- 服务器具有**固定公网 IP**

---

## DNS 解析配置

要让其他邮件服务器能将邮件投递到你的服务器，需要在你的域名 DNS 控制台完成以下两步配置。假设你的服务器 IP 为 `1.2.3.4`，想使用子域名 `mail.example.com` 作为邮件域。

### 第一步：添加 A 记录

将子域名指向你的服务器 IP：

| 记录类型 | 主机名 | 值（目标） | TTL |
|---------|--------|-----------|-----|
| A | mail | 1.2.3.4 | 600 |

> 如果使用根域名（`example.com`），主机名填 `@`。

### 第二步：添加 MX 记录

告知其他邮件服务器，发往 `@mail.example.com` 的邮件应投递到哪台主机：

| 记录类型 | 主机名 | 邮件服务器（值） | 优先级 | TTL |
|---------|--------|----------------|--------|-----|
| MX | mail | mail.example.com | 10 | 600 |

> **注意**：MX 记录的值必须是一个主机名（即上一步配置了 A 记录的域名），不能直接填 IP 地址。

配置完成后，可用以下命令验证 DNS 是否生效（通常需要几分钟到数小时）：

```bash
# 查询 A 记录
dig A mail.example.com

# 查询 MX 记录
dig MX mail.example.com
```

---

## 构建与运行

```bash
# 克隆并构建
git clone https://github.com/Jlan45/temporaryMailbox.git
cd temporaryMailbox
go build -o temporaryMailbox .

# 单域名模式（将 mail.example.com 替换为你的实际域名）
SUBDOMAIN=mail.example.com ./temporaryMailbox

# 多域名模式（逗号分隔，用户可在 Web UI 中自行选择域名）
SUBDOMAINS=mail.example.com,mail2.example.com,inbox.example.org ./temporaryMailbox
```

> **说明**：`SUBDOMAINS` 优先级高于 `SUBDOMAIN`。设置 `SUBDOMAINS` 后，`SUBDOMAIN` 将被忽略。

程序启动后：
- HTTP 服务监听 `:80`，提供 Web UI 和 REST 接口
- SMTP 服务监听 `:25`，接收来自其他邮件服务器的投递
- 如果配置了多个域名，用户可以在 Web 页面的下拉列表中选择想要使用的域名

---

## API 接口

### 获取可用域名列表

```
GET /getDomains
```

**响应示例：**
```json
{
    "domains": ["mail.example.com", "mail2.example.com"]
}
```

---

### 获取随机邮件地址

```
GET /getAddress?domain=mail.example.com
```

> `domain` 参数可选。如果未提供，默认使用第一个配置的域名。

**响应示例：**
```json
{
    "address": "79lb91nk@mail.example.com",
    "domain": "@mail.example.com",
    "random": "79lb91nk"
}
```

---

### 获取邮件列表

```
GET /getMailList/:random
```

**响应示例：**
```json
{
    "mails": [
        {
            "from": "sender@example.com",
            "title": "SMTP 邮件测试"
        }
    ]
}
```

---

### 获取并弹出最新邮件

```
GET /getMail/:random
```

**响应示例：**
```json
{
    "mail": {
        "content": "邮件发送测试...",
        "from": "sender@example.com",
        "title": "SMTP 邮件测试"
    }
}
```

> 调用 `/getMail` 后，对应邮件会从内存中移除（pop），无需担心邮件堆积占满内存。

---

## Web UI

直接访问 `http://<你的域名>/` 即可使用内置的网页界面，支持一键生成地址、查看邮件列表和读取邮件内容。

演示地址：http://mail.nothinglikethis.asia/
