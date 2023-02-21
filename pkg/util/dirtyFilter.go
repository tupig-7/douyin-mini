package util

import (
	"github.com/importcjj/sensitive"
)

func New() *sensitive.Filter {
	filter := sensitive.New()
	err := filter.LoadNetWordDict("https://douyin-mini12306.oss-cn-beijing.aliyuncs.com/dict.txt?Expires=1676973360&OSSAccessKeyId=TMP.3Kfs9ZoHvhauiysAyYXG1VSptyKb6aZrW9qC5Bvn4LPvpY5x9ezrPW68iFvQEQpFc2tCQTBA4nhJKPK8TAxADyzf4f2uVi&Signature=QYv2Kdteim1VI5Fl08R7li7QMDw%3D")
	if err != nil {
		return nil
	}
	return filter
}

// Filter 敏感词过滤
func Filter(dirty string) string {
	f := New()
	clean := f.Replace(dirty, '*')
	return clean
}
