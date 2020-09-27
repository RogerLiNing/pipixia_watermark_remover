package pipixia

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonData struct {
	Data Data `json:"data"`
}

type Data struct {
	Item Item `json:"item"`
}

type Item struct {
	OriginVideoURL OriginVideoURL `json:"origin_video_download"`
	OriginVideoId  string         `json:"origin_video_id"`
}

type OriginVideoURL struct {
	UrlList []URL `json:"url_list"`
}

type URL struct {
	URL string `json:"url"`
}

func getVideoLink(id string) (string, error) {

	client := &http.Client{}
	// 通过这个接口获取视频信息，其中包括带有水印的链接
	url := "https://h5.pipix.com/bds/webapi/item/detail/?item_id=" + id

	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	jsonByteData := []byte(string(body))

	jsonData := JsonData{}
	json.Unmarshal(jsonByteData, &jsonData)
	var videoLink = ""

	if len(jsonData.Data.Item.OriginVideoURL.UrlList) == 2 {
		videoLink = jsonData.Data.Item.OriginVideoURL.UrlList[0].URL
	} else if len(jsonData.Data.Item.OriginVideoId) > 0 {
		videoLink, _ = getVideoDownloadLink(jsonData.Data.Item.OriginVideoId)
	}

	return videoLink, nil

}
