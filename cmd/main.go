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

	var rows []*cdp.Node
	// var capabilityRows []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Navigate(URL),
		chromedp.Nodes("table.primary-table > tbody", &rows, chromedp.ByQueryAll),
		chromedp.ActionFunc(func(c context.Context) error {
			// depth -1 for the entire subtree
			// do your best to limit the size of the subtree
			for _, n := range rows {
				dom.RequestChildNodes(n.NodeID).WithDepth(-1).Do(c)
			}
			return nil
		}),
	)

	if err != nil {
		log.Fatalln(err)
	}

	// rows[0].Children -> "tr"
	for _, n := range rows[0].Children {
		if isSeriveRow(n.Attributes[1]) || isCapabilityRow((n.Attributes[1])) || isCategoryRow(n.Attributes[1]) {
			// n.Children[0] -> "th"
			// n.Children[0].Children[0] -> "th > a" or "th > th value"
			c := n.Children[0].Children[0]
			if c.ChildNodeCount != 0 {
				// if ChildNodeCount is not equal 0, "th > a"
				// access children node to get anchor value
				c = c.Children[0]
			}
			fmt.Println(strings.TrimSpace(c.NodeValue))
		}
	}
}

func isSeriveRow(s string) bool {
	return s == "service-row"
}

func isCapabilityRow(s string) bool {
	return s == "capability-row"
}

func isCategoryRow(s string) bool {
	return s == "category-row"
}
