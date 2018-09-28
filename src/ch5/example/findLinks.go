package main
/*
	本文件演示递归的初步使用，
	文件中用到了golang.org/x/net/html的包，把该包的文件放到src目录下即可
*/
import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// visit appends to links each link found in n and returns the result.
//函数解析HTML标准输入，通过递归函数visit获得links（链接），并打印出这些links
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}


func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
