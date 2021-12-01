package review

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/communication"
)

const exampleCommandEdit string = "/edit__communication__review " +
	"{\"review_id\": 1, \"user_id\": 1, \"item_id\": 1, \"text\": \"review text\", \"rating\": 5}"

func (c *SubCommander) edit(inputMsg *tgbotapi.Message) {
	senderFunc := createLogTextSenderFunc(c.bot, inputMsg.Chat.ID, "edit")

	args := inputMsg.CommandArguments()

	var review communication.Review
	err := json.Unmarshal([]byte(args), &review)
	if err != nil {
		log.Printf(`wrong args "%s": %v`, args, err)
		senderFunc(fmt.Sprintf("wrong args.\nExample command:\n%s", exampleCommandEdit))
		return
	}

	err = c.reviewService.Update(review.ReviewID, review)
	if err != nil {
		log.Printf(`fail to edit review "%v": %v`, review, err)
		senderFunc("review not found")
		return
	}

	senderFunc(fmt.Sprintf("successfully edited review (id: %d)", review.ReviewID))
}
