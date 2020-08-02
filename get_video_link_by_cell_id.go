package pipixia

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonData2 struct {
	Data Data2 `json:"data"`
}

type Data2 struct {
	Comment Comment `json:"comment"`
}

type Comment struct {
	VideoFallBack VideoFallBack `json:"video_fallback"`
}

type VideoFallBack struct {
	UrlList []URL `json:"url_list"`
}

type URL2 struct {
	URL string `json:"url"`
}


func getVideoLinkByCellId(cellId string, cellType string) (string, error) {

	client := &http.Client{}
	// 通过这个接口获取视频信息，其中包括带有水印的链接
	url := "https://h5.pipix.com/bds/webapi/cell/detail/?cell_id=" + cellId + "&cell_type=" + cellType + "&source=share"

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

	jsonData := JsonData2{}
	json.Unmarshal(body, &jsonData)

	var videoLink = ""

	if len(jsonData.Data.Comment.VideoFallBack.UrlList) == 2 {
		videoLink = jsonData.Data.Comment.VideoFallBack.UrlList[0].URL
	}

	return videoLink, nil

}
