package LineBotTest

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/go_line/Service"
	"testing"
)

func TestAttractionAPI(t *testing.T) {
	attractionDetail := Service.GetTaipeiAttraction(1)

	message := fmt.Sprintf("%s %s", attractionDetail.Name, "Map: "+attractionDetail.URL)

	if message != "" {
		t.Log(message)
	} else {
		t.Log("fail")
	}
}
