package main
/*
	本文件通过递归的方式遍历整个HTML结点树，并输出树的结构。
*/
import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		//每遇到一个HTML元素标签，就将其入栈，并输出
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
		//outline有入栈操作，但没有相对应的出栈操作。当outline调用自身时，被调用者接收的是stack的拷贝。被调用者的入栈操作，修改的是stack的拷贝，而不是调用者的stack,因对当函数
		//返回时,调用者的stack并未被修改。
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {outline(stack, c)}
}
