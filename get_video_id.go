package pipixia

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func getVideoId(url string) (string,string, string, string, error) {
	var postID string
	if strings.Contains(url, "/post/") {

		//解析正则表达式，如果成功返回解释器
		reg := regexp.MustCompile(`(?m)\/post\/(\d*)`)

		//根据规则提取关键信息
		result := reg.FindAllStringSubmatch(url,-1)

		if len(result) > 0 {
			postID = result[0][1]
			return "", "", "", postID, nil
		} else {
			return "", "", "", "", nil
		}

	}

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")

	if err != nil {
		log.Fatal(err)
		return "", "", "", "", err
	}

	var lastUrlQuery string
	var videoId string
	var cellId string
	var cellType string

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {

		if len(via) > 10 {
			return errors.New("too many redirects")
		}
		lastUrlQuery = req.URL.RequestURI()
		cellId = req.URL.Query().Get("cell_id")
		cellType = req.URL.Query().Get("cell_type")

		return nil
	}

	response, err := client.Do(request)


	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		return "", "", "", "", nil
	}


	if len(strings.Split(strings.Split(lastUrlQuery, "?")[0], "/")) > 1 {
		videoId = strings.Split(strings.Split(lastUrlQuery, "?")[0], "/")[2]
	}

	return videoId, cellId, cellType, postID, nil
}
