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
	"os"
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

	jsonErr := json.Unmarshal([]byte(res.RequestBody), &event)
	if jsonErr != nil {
		return events.APIGatewayProxyResponse{}, jsonErr
	}

	userid := fmt.Sprintf("%v", event.Events[0].Source.UserID)
	text := fmt.Sprintf("%v", event.Events[0].Message.Text)

	LineBotInit(userid, text)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil

}

func LineBotInit(userid string, text string) {

	client := &http.Client{}

	bot, err := linebot.New(os.Getenv("channelSecret"), os.Getenv("channelToken"), linebot.WithHTTPClient(client))
	if err != nil {
		log.Println(err)
	}

	PushMessage(userid, bot, text)

}

func PushMessage(userid string, bot *linebot.Client, text string) {

	if text == "night" && userid == os.Getenv("myUserID") {
		if _, err := bot.PushMessage(os.Getenv("yuUserID"), linebot.NewTextMessage("寶寶晚安")).Do(); err != nil {
			log.Println(err)
		}
		if _, err := bot.PushMessage(os.Getenv("yuUserID"), linebot.NewStickerMessage("6362", "11087943")).Do(); err != nil {
			log.Println(err)
		}
		return
	}

	if text == "love" && userid == os.Getenv("myUserID") {
		if _, err := bot.PushMessage(os.Getenv("yuUserID"), linebot.NewTextMessage("寶寶我愛你")).Do(); err != nil {
			log.Println(err)
		}
		if _, err := bot.PushMessage(os.Getenv("yuUserID"), linebot.NewStickerMessage("11538", "51626502")).Do(); err != nil {
			log.Println(err)
		}
		return
	}

	rand.Seed(time.Now().UnixNano())
	recommendationAttraction := GetTaipeiAttraction(rand.Intn(16))

	if _, err := bot.PushMessage(userid, linebot.NewTextMessage(recommendationAttraction)).Do(); err != nil {
		log.Println(err)
	}

}

func GetTaipeiAttraction(page int) string {

	client := &http.Client{}

	var strBuild strings.Builder
	strBuild.WriteString("https://www.travel.taipei/open-api/zh-tw/Attractions/All?page=")
	strBuild.WriteString(strconv.Itoa(page))

	url := strBuild.String()

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

	closeErr := resp.Body.Close()

	if closeErr != nil {
		fmt.Println(err)
		return "Please Wait..."
	}

	if err != nil {
		fmt.Println(err)
		return "Please Wait..."
	}

	var attraction ApiModel.Attractions

	jsonErr := json.Unmarshal(body, &attraction)
	if jsonErr != nil {
		fmt.Println(err)
		return "Please Wait..."
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(31)

	return attraction.Data[randomIndex].Name
}
