package communication

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/communication/review"
	"github.com/ozonmp/omp-bot/internal/app/path"
	serviceReview "github.com/ozonmp/omp-bot/internal/service/communication/review"
)

type subCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type Commander struct {
	bot                *tgbotapi.BotAPI
	reviewSubCommander subCommander
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	reviewService := serviceReview.NewService()
	return &Commander{
		bot:                bot,
		reviewSubCommander: review.NewSubCommander(bot, reviewService),
	}
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "review":
		c.reviewSubCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("communication.Commander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "review":
		c.reviewSubCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("communication.Commander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
