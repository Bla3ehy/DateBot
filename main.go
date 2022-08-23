package main

import (
	"encoding/json"
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/go_line/ApiModel"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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

	json.Unmarshal([]byte(res.RequestBody), &event)

	userid := fmt.Sprintf("%v", event.Events[0].Source.UserID)
	text := fmt.Sprintf("%v", event.Events[0].Message.Text)

	LineBotInit(userid, text)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil

}

func LineBotInit(userid string, text string) {

	client := &http.Client{}

	bot, err := linebot.New(`a1c7ddc08b78297946ab4b7731c19ce5`, `6fwhdTxYdM+4DgaflnZ5PBRzxmx5b3MlNZcBQoYfP0UokaeHwSQHg8eHi6QCXaOD7pAgErYafbhgvBJbtipX7yTqe8LXV28LWu6B+Tg+SDd4yMxXve0wxja/15W7wQPiF8eYorcCPFuebz0nV5H2+QdB04t89/1O/w1cDnyilFU=`, linebot.WithHTTPClient(client))
	if err != nil {
		log.Println(err)
	}

	rand.Seed(time.Now().UnixNano())

	recommendationAttraction := GetTaipeiAttraction(rand.Intn(16))

	PushMessage(userid, bot, recommendationAttraction)

}

func PushMessage(userid string, bot *linebot.Client, recommendationAttraction string) {

	if _, err := bot.PushMessage(userid, linebot.NewTextMessage(recommendationAttraction)).Do(); err != nil {
		log.Println(err)
	}

	if userid == `U974a81e8995bb6d231cc06a749e7ddbf` {
		if _, err := bot.PushMessage(userid, linebot.NewTextMessage("寶寶我愛妳喔")).Do(); err != nil {
			log.Println(err)
		}

		if _, err := bot.PushMessage(userid, linebot.NewStickerMessage("11538", "51626502")).Do(); err != nil {
			log.Println(err)
		}
	}
}

func GetTaipeiAttraction(page int) string {

	client := &http.Client{}

	var str_build strings.Builder
	str_build.WriteString("https://www.travel.taipei/open-api/zh-tw/Attractions/All?page=")
	str_build.WriteString(strconv.Itoa(page))

	url := str_build.String()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return "Please Wait..."
	}

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "Please Wait..."
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "Please Wait..."
	}
	var attraction ApiModel.Attractions

	json.Unmarshal(body, &attraction)

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(31)

	return attraction.Data[randomIndex].Name
}
