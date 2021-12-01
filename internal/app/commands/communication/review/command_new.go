package review

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/communication"
)

const exampleCommandNew string = "/new__communication__review " +
	"{\"user_id\": 1, \"item_id\": 1, \"text\": \"review text\", \"rating\": 5}"

func (c *SubCommander) new(inputMsg *tgbotapi.Message) {
	senderFunc := createLogTextSenderFunc(c.bot, inputMsg.Chat.ID, "new")

	args := inputMsg.CommandArguments()

	var review communication.Review
	err := json.Unmarshal([]byte(args), &review)
	if err != nil {
		log.Printf(`wrong args "%s": %v`, args, err)
		senderFunc(fmt.Sprintf("wrong args.\nExample command:\n%s", exampleCommandNew))
		return
	}

	reviewID, err := c.reviewService.Create(review)
	if err != nil {
		log.Printf(`fail to create review "%v": %v`, review, err)
		senderFunc("server error")
		return
	}

	senderFunc(fmt.Sprintf("created new review (id: %d)", reviewID))
}
