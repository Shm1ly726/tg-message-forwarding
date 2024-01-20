# tg-message-forwarding
## 简介
顾名思义，是一个tg转发消息的小工具，在写自动化的时候，需要监听tg群组里的osmedeus漏洞推送同步转发到钉钉上，需要就造了这么一个轮子(基于Golang)

能力有限，代码写得很垃圾，各位师傅们轻喷，也欢迎各位师傅指导与讨论


## 主要功能

1. 监听tg群组里的所有消息并且转发到钉钉上
## 使用说明
依赖库：
- github.com/blinkbean/dingtalk
- github.com/go-telegram-bot-api/telegram-bot-api

注册一个专门监听群组消息的tg机器人并且记录下access token
```
@BotFather
osm_bot
xxx_osm_bot
```
![image](https://github.com/Shm1ly726/tg-message-forwarding/assets/142187061/fa0b4db4-1317-4d07-be1e-da235a7b46ea)

创建TG频道后拉监听机器人进去后，需要修改脚本四处位置：
```go
var dingToken = "钉钉机器人Webhook TOKEN"
var dingSecret = "钉钉机器人加签Secret"
bot, err := tgbotapi.NewBotAPI("tg机器人telegram_api_token")
if update.ChannelPost.Chat.UserName == "需要监听的群组名称
```

编译运行
```go
go mod init
go mod tidy
go build
```
## 运行截图
后台挂载监听
![image](https://github.com/Shm1ly726/tg-message-forwarding/assets/142187061/a622e62a-c474-4c6c-b7cd-4d6947aadbc2)

使用osmedeus测试推送转发
![image](https://github.com/Shm1ly726/tg-message-forwarding/assets/142187061/482940ca-075e-4d28-ac94-71af7e966013)

