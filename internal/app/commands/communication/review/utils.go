package review

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func createLogTextSenderFunc(bot *tgbotapi.BotAPI, chatID int64, command string) func(text string) {
	return func(text string) {
		msg := tgbotapi.NewMessage(chatID, text)
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("review.SubCommander.%s: error sending reply message to chat - %v", command, err)
		}
	}
}

func createLogMsgSenderFunc(bot *tgbotapi.BotAPI, command string) func(msg tgbotapi.MessageConfig) {
	return func(msg tgbotapi.MessageConfig) {
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("review.SubCommander.%s: error sending reply message to chat - %v", command, err)
		}
	}
}
