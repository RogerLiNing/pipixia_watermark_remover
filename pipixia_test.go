package pipixia

import (
	"testing"
)

func TestGetAvailablePipixiaLink(t *testing.T) {
	t.Log("测试正常的皮皮虾视频链接")
	url := "https://h5.pipix.com/s/JjaQjEj/" // https://h5.pipix.com/s/JDh49ac/ https://h5.pipix.com/s/JjaQjEj/  https://h5.ippzone.com/pp/post/363597059443?zy_to=copy_link
	t.Log(url)
	videoLink, err := WatermarkRemover(url)

	if err != nil {
		t.Fail()
		t.Log("返回视频链接，测试通过")
	}
	if len(videoLink) == 0 {
		t.Fail()
	}
	t.Log(videoLink)
}

func TestGetUnAvailablePipixiaLink(t *testing.T) {
	t.Log("测试正常的皮皮虾视频链接")
	url := "https://h5.pipix.com/s/J200w4Qr0157x/"
	t.Log(url)
	videoLink, err := WatermarkRemover(url)

	if err != nil {
		t.Fail()
	}
	if len(videoLink) == 0 {
		t.Log("返回为空，测试通过")
	}
	t.Log(videoLink)
}
