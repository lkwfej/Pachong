package parse

import "regexp"

func Biaoti(html string) string {
	bt := regexp.MustCompile(`(?i)<title>(.*?)</title>`)
	match := bt.FindStringSubmatch(html)
	if len(match) >= 2 {
		return match[1]
	}
	return "没有标题"
}
