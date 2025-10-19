package main

import (
	"Pachong/fetch"
	"Pachong/parse"
	"Pachong/storage"
	"fmt"
)

type PageResult struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func main() {
	urls := []string{
		"https://www.baidu.com",
	}

	results := make([]PageResult, 0) //建立一个可以自动追加的切片，初始容量为0
	for _, url := range urls {
		u, err := fetch.Gethtml(url)
		if err != nil {
			fmt.Println("打印失败", err)
			continue
		}
		fmt.Println("", u)

		title := parse.Biaoti(u)
		fmt.Printf("%s", title)
		results = append(results, PageResult{URL: url, Title: title})
	}
	storage.Save("result.json", results)
}
