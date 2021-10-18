package review

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const exampleCommandDelete string = "/delete__communication__review 1"

func (c *CommunicationReviewCommander) Delete(inputMsg *tgbotapi.Message) {
	senderFunc := createLogTextSenderFunc(c.bot, inputMsg.Chat.ID, "Delete")

	args := inputMsg.CommandArguments()
	reviewID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		senderFunc(fmt.Sprintf("wrong args.\nExample command:\n%s", exampleCommandDelete))
		return
	}

	ok, err := c.reviewService.Remove(reviewID)
	if err != nil {
		log.Printf("fail remove review (id: %d): %v", reviewID, err)
		senderFunc("server error")
		return
	}
	if !ok {
		log.Printf("fail remove review (id: %d): already removed", reviewID)
		senderFunc("review not found")
		return
	}

	senderFunc(fmt.Sprintf("review (id: %d) successfully removed", reviewID))
}
