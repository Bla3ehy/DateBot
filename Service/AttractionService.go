package Service

import (
	"encoding/json"
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/go_line/ApiModel"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetTaipeiAttraction(page int) ApiModel.AttractionDetail {

	client := &http.Client{}

	var strBuild strings.Builder
	strBuild.WriteString("https://www.travel.taipei/open-api/zh-tw/Attractions/All?page=")
	strBuild.WriteString(strconv.Itoa(page))

	var attractionData ApiModel.AttractionDetail

	url := strBuild.String()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return attractionData
	}

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return attractionData
	}

	body, err := ioutil.ReadAll(resp.Body)

	closeErr := resp.Body.Close()

	if closeErr != nil {
		fmt.Println(err)
		return attractionData
	}

	if err != nil {
		fmt.Println(err)
		return attractionData
	}

	var attraction ApiModel.Attractions

	jsonErr := json.Unmarshal(body, &attraction)
	if jsonErr != nil {
		fmt.Println(err)
		return attractionData
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(31)
	attractionData = attraction.Data[randomIndex]

	return attractionData

}
