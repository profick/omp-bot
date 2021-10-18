package review

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/communication/review"
)

type ReviewCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented

	CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
}

type CommunicationReviewCommander struct {
	bot           *tgbotapi.BotAPI
	reviewService service.ReviewService
}

func NewCommunicationReviewCommander(bot *tgbotapi.BotAPI, service service.ReviewService) *CommunicationReviewCommander {
	return &CommunicationReviewCommander{
		bot:           bot,
		reviewService: service,
	}
}

func (c *CommunicationReviewCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CommunicationReviewCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CommunicationReviewCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Help(msg)
	}
}
