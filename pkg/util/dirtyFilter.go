package util

import (
	"github.com/importcjj/sensitive"
)

func New() *sensitive.Filter {
	filter := sensitive.New()
	err := filter.LoadNetWordDict("https://douyin-mini12306.oss-cn-beijing.aliyuncs.com/dict.txt")
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
