package pipixia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)



func getVideoLinkByPostId(postId string) (string, error) {

	client := &http.Client{}
	// 通过这个接口获取视频信息，其中包括带有水印的链接
	url := "https://h5.ippzone.com/ppapi/share/fetch_content"

	payload := strings.NewReader(`{"pid":` + postId + `,"mid":null,"type":"post"}`)

	request, err := http.NewRequest("POST", url, payload)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)

	}

	// Declared an empty interface of type Array
	var results map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(body, &results)


	data, _ := results["data"].(map[string]interface{})
	post, _ := data["post"].(map[string]interface{})
	var videoLink string
	for _, video := range post["videos"].(map[string]interface{}) {
		videoLink = fmt.Sprintf("%v", video.(map[string]interface{})["url"])
	}

	return videoLink, nil
}
