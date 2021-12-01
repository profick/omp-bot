package review

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/communication"
)

const reviewsPerPage = 5

func (c *SubCommander) list(inputMsg *tgbotapi.Message) {
	senderFunc := createLogTextSenderFunc(c.bot, inputMsg.Chat.ID, "list")
	msgSenderFunc := createLogMsgSenderFunc(c.bot, "list")

	reviews, err := c.reviewService.List(0, reviewsPerPage)
	if err != nil {
		log.Printf("fail retrieving list of reviews: %v", err)
		senderFunc("server error")
		return
	}
	if len(reviews) == 0 {
		reviews = []communication.Review{}
	}

	msgData, err := json.MarshalIndent(reviews, "", "  ")
	if err != nil {
		log.Printf("fail to marshal reviews %v, %v", reviews, err)
		senderFunc("server error")
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, string(msgData))

	offset := reviewsPerPage
	buttonName := "next"
	if len(reviews) < reviewsPerPage {
		// Use `update` button instead of `next`
		offset = 0
		buttonName = "update"
	}
	buttonData, err := createCallbackListPathString(offset)
	if err != nil {
		log.Printf("failed creating callback list path string: %v", err)
		senderFunc("server error")
		return
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttonName, buttonData),
		),
	)
	msgSenderFunc(msg)
}

func createCallbackListPathString(offset int) (string, error) {
	nextCallbackData, err := json.Marshal(CallbackListData{Offset: offset})
	if err != nil {
		return "", err
	}
	callbackPath := path.CallbackPath{
		Domain:       "communication",
		Subdomain:    "review",
		CallbackName: "list",
		CallbackData: string(nextCallbackData),
	}
	return callbackPath.String(), nil
}
