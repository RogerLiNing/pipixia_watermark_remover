package pipixia

func WatermarkRemover(url string) (string, error) {
	videoId, err := getVideoId(url)

	videoLink, err := getVideoLink(videoId)

	return videoLink, err
}
