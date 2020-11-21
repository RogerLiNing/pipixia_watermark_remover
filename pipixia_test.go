package pipixia

import (
	"testing"
)

func TestGetAvailablePipixiaLink(t *testing.T) {
	t.Log("测试正常的皮皮虾视频链接")
	// VideoId https://h5.pipix.com/s/Jx8aW6Y/
	// PostID  https://h5.ippzone.com/pp/post/363597059443
	// CellId  https://h5.pipix.com/s/Jum16RV/
	url := "https://h5.pipix.com/s/Jum16RV/"

	videoLink, err := WatermarkRemover(url)

	if err != nil {
		t.Fail()
	}

	if len(videoLink) > 0 {
		t.Log("测试通过")
	}

}
