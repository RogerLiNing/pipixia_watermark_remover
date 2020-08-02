package pipixia

func WatermarkRemover(url string) (string, error) {
	var videoLink string
	videoId,cellId, cellType, err := getVideoId(url)

	if len(cellId) > 0 {
		videoLink, err = getVideoLinkByCellId(cellId,cellType)
	} else {
		videoLink, err = getVideoLink(videoId)
	}

	return videoLink, err
}
