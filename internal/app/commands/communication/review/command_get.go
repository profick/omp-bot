package review

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const exampleCommandGet string = "/get__communication__review 1"

func (c *SubCommander) get(inputMsg *tgbotapi.Message) {
	senderFunc := createLogTextSenderFunc(c.bot, inputMsg.Chat.ID, "get")

	args := inputMsg.CommandArguments()
	reviewID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		senderFunc(fmt.Sprintf("wrong args.\nExample command:\n%s", exampleCommandGet))
		return
	}

	review, err := c.reviewService.Describe(reviewID)
	if err != nil {
		log.Printf("fail to get review with id %d: %v", reviewID, err)
		senderFunc("review not found")
		return
	}

	msgData, err := json.MarshalIndent(review, "", "  ")
	if err != nil {
		log.Printf("fail to marshal review %s, %v", review, err)
		senderFunc("server error")
		return
	}
	senderFunc(string(msgData))
}
