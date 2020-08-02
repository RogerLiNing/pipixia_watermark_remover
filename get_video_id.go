package pipixia

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getVideoId(url string) (string, string, string, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")

	if err != nil {
		log.Fatal(err)
		return "","","", err
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
		return "", "", "", err
	}

	if len(strings.Split(strings.Split(lastUrlQuery, "?")[0], "/")) > 1 {
		videoId = strings.Split(strings.Split(lastUrlQuery, "?")[0], "/")[2]
	}

	return videoId, cellId, cellType, nil
}
