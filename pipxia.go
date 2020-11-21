package pipixia

import "fmt"

func WatermarkRemover(url string) (string, error) {
	var videoLink string
	videoId,cellId, cellType, postID, err := getVideoId(url)

	if len(cellId) > 0 {
		fmt.Println("cellId:" + cellId)
		videoLink, err = getVideoLinkByCellId(cellId,cellType)
	} else if len(postID) > 0 {
		fmt.Println("postId:" + postID)
		videoLink, err = getVideoLinkByPostId(postID)
	} else {
		fmt.Println("videoId:" + videoId)
		videoLink, err = getVideoLink(videoId)
	}

	return videoLink, err
}
