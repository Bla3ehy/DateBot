package main

import (
	"encoding/json"
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/go_line/ApiModel"
	"github.com/YOUR-USER-OR-ORG-NAME/go_line/Service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var event ApiModel.Event

	body := request.Body
	res := ApiModel.Response{
		RequestBody: body,
	}

	jsonErr := json.Unmarshal([]byte(res.RequestBody), &event)
	if jsonErr != nil {
		return events.APIGatewayProxyResponse{}, jsonErr
	}

	LineBotInit(event)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil

}

func LineBotInit(event ApiModel.Event) {

	client := &http.Client{}

	bot, err := linebot.New(os.Getenv("channelSecret"), os.Getenv("channelToken"), linebot.WithHTTPClient(client))
	if err != nil {
		log.Println(err)
	}

	userid := fmt.Sprintf("%v", event.Events[0].Source.UserID)
	text := fmt.Sprintf("%v", event.Events[0].Message.Text)

	rand.Seed(time.Now().UnixNano())
	attractionData := Service.GetTaipeiAttraction(rand.Intn(16))

	if attractionData.Name == "" {
		log.Println("No Data")
		return
	}

	PushMessage(userid, bot, text, attractionData)

}

func PushMessage(userid string, bot *linebot.Client, text string, attractionDetail ApiModel.AttractionDetail) {

	if userid != os.Getenv("yuUserID") {
		if text == "night" {
			if _, err := bot.PushMessage(os.Getenv("yuUserID"), linebot.NewTextMessage("寶寶晚安")).Do(); err != nil {
				log.Println(err)
			}
			if _, err := bot.PushMessage(os.Getenv("yuUserID"), linebot.NewStickerMessage("6362", "11087943")).Do(); err != nil {
				log.Println(err)
			}
			return
		}

		if text == "love" {
			if _, err := bot.PushMessage(os.Getenv("yuUserID"), linebot.NewTextMessage("寶寶我愛你")).Do(); err != nil {
				log.Println(err)
			}
			if _, err := bot.PushMessage(os.Getenv("yuUserID"), linebot.NewStickerMessage("11538", "51626502")).Do(); err != nil {
				log.Println(err)
			}
			return
		}
	}

	message := fmt.Sprintf("%s  %s", attractionDetail.Name, "Map: "+attractionDetail.URL)

	if _, err := bot.PushMessage(userid, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Println(err)
	}

}
