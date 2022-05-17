package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const URL string = "https://azure.microsoft.com/ja-jp/global-infrastructure/services/?regions=all&products=all"

func main() {
	doMain()
}

func doMain() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var nodes []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Navigate(URL),
		//chromedp.Nodes(`document.querySelector("#primary-table > tbody > tr:nth-child(7) > th > a")`, &nodes, chromedp.ByJSPath)
		chromedp.Nodes(`document.querySelector("#primary-table > tbody > tr:nth-child(8) > th")`, &nodes, chromedp.ByJSPath),
	)

	if err != nil {
		log.Fatalln(err)
	}

	for _, n := range nodes {
		fmt.Println(n.Children[0].NodeValue)
	}
}
