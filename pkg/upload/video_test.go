package upload

import (
	"testing"
)

func TestExactCoverFromVideo(t *testing.T) {
	pathVideo := "D:/STUDY/Project/douyinMini/douyin-mini/storage/uploads/53/61fac9094c9793208b20c3f0e0541ff1.mp4"
	pathImg := "D:/STUDY/Project/douyinMini/douyin-mini/storage/uploads/53/61fac9094c9793208b20c3f0e0541ff1.png"
	err := ExactCoverFromVideo(pathVideo, pathImg)
	if err != nil {
		t.Error(err)
	}
}
