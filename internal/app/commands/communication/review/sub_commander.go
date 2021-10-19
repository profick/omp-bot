package review

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/communication"
)

type SubCommanderService interface {
	Describe(reviewID uint64) (*communication.Review, error)
	List(cursor uint64, limit uint64) ([]communication.Review, error)
	Create(communication.Review) (uint64, error)
	Update(reviewID uint64, review communication.Review) error
	Remove(reviewID uint64) (bool, error)
}

type SubCommander struct {
	bot           *tgbotapi.BotAPI
	reviewService SubCommanderService
}

func NewSubCommander(bot *tgbotapi.BotAPI, service SubCommanderService) *SubCommander {
	return &SubCommander{
		bot:           bot,
		reviewService: service,
	}
}

func (c *SubCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.callbackList(callback, callbackPath)
	default:
		log.Printf("review.SubCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *SubCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "get":
		c.get(msg)
	case "list":
		c.list(msg)
	case "delete":
		c.delete(msg)
	case "new":
		c.new(msg)
	case "edit":
		c.edit(msg)
	default:
		c.help(msg)
	}
}
