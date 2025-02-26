package test

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"internal_api/pkg/sqllite_base"
	"internal_api/pkg/websocket_base"
	"log"
	"testing"
)

func TestH(t *testing.T) {
	fmt.Println(111)
}

func TestFeed(t *testing.T) {
	// 要订阅的 RSS Feed 地址
	rssFeedURL := "https://rsshub.app/sina/csj"

	// 创建一个新的 RSS 解析器
	fp := gofeed.NewParser()

	// 解析 RSS Feed
	feed, err := fp.ParseURL(rssFeedURL)
	if err != nil {
		log.Fatal("Error parsing RSS Feed:", err)
	}

	// 打印 Feed 标题
	fmt.Println("Feed Title:", feed.Title)

	// 遍历所有的 Feed 项目并输出
	for _, item := range feed.Items {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Published: %s\n", item.Published)
		fmt.Println("Description:", item.Description)
		fmt.Println("---------------")
	}
}

func TestSqllite(t *testing.T) {
	sqllite_base.Init()
}

func TestSocket(t *testing.T) {
	websocket_base.Init()
}
