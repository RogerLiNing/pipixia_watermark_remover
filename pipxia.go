package pipixia

func WatermarkRemover(url string) (string, error) {
	var videoLink string
	videoId,cellId, cellType, postID, err := getVideoId(url)

	if len(cellId) > 0 {
		videoLink, err = getVideoLinkByCellId(cellId,cellType)
	} else if len(postID) > 0 {
		videoLink, err = getVideoLinkByPostId(postID)
	} else {
		videoLink, err = getVideoLink(videoId)
	}

	return videoLink, err
}
