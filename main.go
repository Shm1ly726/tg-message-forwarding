package main

import (
    "github.com/blinkbean/dingtalk"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
)

var msg string

type DingTalkMessage struct {
    MsgType string                 `json:"msgtype"`
    Text    map[string]interface{} `json:"text"`
}

func dingTalkBot(tg_massage string) {
    var dingToken = "钉钉机器人Webhook TOKEN"
    var dingSecret = "钉钉机器人加签Secret"
    cli := dingtalk.InitDingTalkWithSecret(dingToken, dingSecret)
    cli.SendTextMessage(tg_massage)
}

func tg_Bot() {
    bot, err := tgbotapi.NewBotAPI("tg机器人telegram_api_token")
    if err != nil {
        log.Panic(err)
    }
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates, err := bot.GetUpdatesChan(u)
    for update := range updates {
        if update.ChannelPost == nil {
            continue
        }
        if update.ChannelPost.Chat.UserName == "需要监听的群组名称" {
            msg = update.ChannelPost.Text
            dingTalkBot(msg)
        }
    }
}

func main() {
    tg_Bot()
}
