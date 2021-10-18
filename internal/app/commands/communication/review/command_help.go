package review

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const helpMessage string = "/help__communication__review - print help\n" +
	"/get__communication__review <REVIEW_ID> - print review information in json format\n" +
	"/list__communication__review - print list of reviews in json format with pagination\n" +
	"/delete__communication__review <REVIEW_ID> - delete a review\n" +
	"/new__communication__review <REVIEW_DATA> - create new review. <REVIEW_DATA> example:" +
	"{\"user_id\": 1, \"item_id\": 1, \"text\": \"review text\", \"rating\": 5}\n" +
	"/edit__communication__review <EDIT_REVIEW_DATA> - edit existing review. <EDIT_REVIEW_DATA> example:" +
	"{\"review_id\": 1, \"user_id\": 1, \"item_id\": 1, \"text\": \"review text\", \"rating\": 5}\n"

func (c *CommunicationReviewCommander) Help(inputMsg *tgbotapi.Message) {
	senderFunc := createLogTextSenderFunc(c.bot, inputMsg.Chat.ID, "Help")
	senderFunc(helpMessage)
}
