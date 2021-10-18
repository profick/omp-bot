package communication

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/communication/review"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/communication/review"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CommunicationCommander struct {
	bot             *tgbotapi.BotAPI
	reviewCommander Commander
}

func NewCommunicationCommander(bot *tgbotapi.BotAPI) *CommunicationCommander {
	reviewService := service.NewCommunicationReviewService()
	return &CommunicationCommander{
		bot:             bot,
		reviewCommander: review.NewCommunicationReviewCommander(bot, reviewService),
	}
}

func (c *CommunicationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "review":
		c.reviewCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CommunicationCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *CommunicationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "review":
		c.reviewCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CommunicationCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
