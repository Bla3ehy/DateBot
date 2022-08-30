package Service

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/go_line/ApiModel"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"os"
)

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
