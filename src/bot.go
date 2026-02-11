package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var bot *telego.Bot
var message []string

func botHandler() {
	ctx := context.Background()
	
	var err error
	bot, err = telego.NewBot(cfg.Token, telego.WithDefaultDebugLogger())
	if err != nil {logger.Error("Failed to create the bot", "err", err.Error()); os.Exit(1)}

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)
	bh, _ := th.NewBotHandler(bot, updates)
	logger.Debug("BotHandler create")
	bh.Handle(comHandler, th.AnyCommand())
	defer func() { _ = bh.Stop() }()

	_ = bh.Start()
	logger.Debug("BotHandler start")

}

func pluginCheck(update telego.Update) bool {
	if update.Message == nil || len(update.Message.Text) == 0 { return false }

	if update.Message.Text[0] == '/' {
		update.Message.Text = update.Message.Text[1:] }

	pluginsListMaker() //for dynamic checking of plugins
	message = strings.Fields(update.Message.Text)
	if _, ok := plst[message[0]]; ok { 
		logger.Debug("Plugin found")
		return true 
	}

	return false
}

func comHandler(ctx *th.Context, update telego.Update) error {
	if update.Message.Text == "/start" { comStart(ctx, update) 

	} else if !pluginCheck(update) {
		bot.SendMessage(ctx, tu.Messagef(
			tu.ID(update.Message.Chat.ID), "Plugin not found\nPlugins available: %v", plst))
		return errors.New("Plugin not found")

	} else if !permissionCheck(update.Message.From.Username) {
		bot.SendMessage(ctx, tu.Messagef(
			tu.ID(update.Message.Chat.ID), "Permission denied"))

	} else { 
		msg, _ := bot.SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID), "Execute"))
		msgID := msg.MessageID
		messages := pluginExecute() 

		if len(messages) != 0 {
			bot.EditMessageText(ctx, &telego.EditMessageTextParams{
				ChatID:    tu.ID(update.Message.Chat.ID),
				MessageID: msgID,
				Text:      "Report sending...", })

			for count, msg := range messages { 
				messageHandler(msg, ctx, update) 
				bot.EditMessageText(ctx, &telego.EditMessageTextParams{
				ChatID:    tu.ID(update.Message.Chat.ID),
				MessageID: msgID,
				Text:      counter(count, len(messages)),
		})
			}
		}
		bot.DeleteMessage(ctx, &telego.DeleteMessageParams{
			ChatID:		tu.ID(update.Message.Chat.ID),
			MessageID:	msgID,
		})
	}

	return nil
}

func permissionCheck(user string) bool {
	if slices.Contains(cfg.Users, user) { 
		logger.Debug("Permission allowed", "user", user)
		return true 
	} else if slices.Contains(cfg.AllowsPlugins, message[0]) {
		logger.Debug("Permission allowed", "user", user)
		return true 
	} else {
	logger.Debug("Permission denied", "user", user)
	return false
	}
}

func comStart(ctx *th.Context, update telego.Update) error {
		bot.SendMessage(ctx, tu.Messagef(
			tu.ID(update.Message.Chat.ID),
			"This is a bot for controlling a computer using a telegram bot, you can find it at the link below\ngithub.com/Konare1ka/TG-commander",
		))
		return nil
}

func messageHandler(msg string, ctx *th.Context, update telego.Update) {
	logger.Debug(msg)
	switch msg[:3] {
	case "img":
		bot.SendPhoto(ctx, tu.Photo(update.Message.Chat.ChatID(), tu.File(mustOpen(msg[4:]))))
	case "vid":
		bot.SendVideo(ctx, tu.Video(update.Message.Chat.ChatID(), tu.File(mustOpen(msg[4:]))))
	case "aud":
		bot.SendAudio(ctx, tu.Audio(update.Message.Chat.ChatID(), tu.File(mustOpen(msg[4:]))))
	case "doc":
		bot.SendDocument(ctx, tu.Document(update.Message.Chat.ChatID(), tu.File(mustOpen(msg[4:]))))
	default:
		bot.SendMessage(ctx, tu.Message(tu.ID(update.Message.Chat.ID), msg))
	}
}

func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil { logger.Error("Can't open file", "file", filename) }
	return file
}

func counter(count int, lenght int) string {
	return fmt.Sprintf("%d/%d", count + 1, lenght)
} 