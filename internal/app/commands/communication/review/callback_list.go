package review

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/communication"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *CommunicationReviewCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	senderFunc := createLogTextSenderFunc(c.bot, callback.Message.Chat.ID, "CallbackList")
	msgSenderFunc := createLogMsgSenderFunc(c.bot, "CallbackList")

	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("CommunicationReviewCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		senderFunc("server error")
		return
	}

	reviews, err := c.reviewService.List(uint64(parsedData.Offset), reviewsPerPage)
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
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, string(msgData))

	inlineButtons, err := callbackListInlineButtons(parsedData.Offset, len(reviews))
	if err != nil {
		log.Printf("failed creating inline buttons: %v", err)
		senderFunc("server error")
		return
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(inlineButtons...),
	)
	msgSenderFunc(msg)
}

func callbackListInlineButtons(currentOffset, reviewsFound int) ([]tgbotapi.InlineKeyboardButton, error) {
	var inlineButtons []tgbotapi.InlineKeyboardButton
	if currentOffset != 0 {
		prevOffset := currentOffset - reviewsPerPage
		if prevOffset < 0 {
			prevOffset = 0
		}
		prevButtonData, err := createCallbackListPathString(prevOffset)
		if err != nil {
			return nil, err
		}
		inlineButtons = append(inlineButtons,
			tgbotapi.NewInlineKeyboardButtonData("prev", prevButtonData))
	}

	// Last page. Change `next` button to `update` button
	if reviewsFound < reviewsPerPage {
		updateButtonData, err := createCallbackListPathString(currentOffset)
		if err != nil {
			return nil, err
		}
		inlineButtons = append(inlineButtons,
			tgbotapi.NewInlineKeyboardButtonData("update", updateButtonData))
	} else {
		nextOffset := currentOffset + reviewsPerPage
		nextButtonData, err := createCallbackListPathString(nextOffset)
		if err != nil {
			return nil, err
		}
		inlineButtons = append(inlineButtons,
			tgbotapi.NewInlineKeyboardButtonData("next", nextButtonData))
	}
	return inlineButtons, nil
}
