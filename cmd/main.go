package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

const URL string = "https://azure.microsoft.com/ja-jp/global-infrastructure/services/?regions=all&products=all"

func main() {
	doMain()
}

func doMain() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var serviceRows []*cdp.Node
	var capabilityRows []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Navigate(URL),
		chromedp.Nodes("table.primary-table > tbody > tr.service-row > th", &serviceRows, chromedp.ByQueryAll),
		chromedp.Nodes("table.primary-table > tbody > tr.capability-row > th", &capabilityRows, chromedp.ByQueryAll),
		chromedp.ActionFunc(func(c context.Context) error {
			// depth -1 for the entire subtree
			// do your best to limit the size of the subtree
			for _, n := range serviceRows {
				dom.RequestChildNodes(n.NodeID).WithDepth(-1).Do(c)
			}
			for _, n := range capabilityRows {
				dom.RequestChildNodes(n.NodeID).WithDepth(-1).Do(c)
			}
			return nil
		}),
	)

	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println("=====" + "service-row" + "=====")
	// fmt.Println(len(serviceRows))
	for _, n := range serviceRows {
		child := n.Children[0]
		if child.ChildNodeCount != 0 {
			child = child.Children[0]
		}
		fmt.Println(strings.TrimSpace(child.NodeValue))
	}

	// fmt.Println("=====" + "capability-row" + "=====")
	// fmt.Println(len(capabilityRows))
	for _, n := range capabilityRows {
		child := n.Children[0]
		if child.ChildNodeCount != 0 {
			child = child.Children[0]
		}
		fmt.Println(strings.TrimSpace(child.NodeValue))
	}
}
