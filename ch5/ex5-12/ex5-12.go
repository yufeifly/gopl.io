package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, v := range os.Args[1:] {
		outline(v)
	}
}
func outline(url string) (string, error){
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	doc, _ := html.Parse(resp.Body)
	//使用匿名函数实现
	var depth int
	var startElement func(n *html.Node)
	var endElement func(n *html.Node)

	startElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	endElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	//1.使用函数值
	forEachNode(doc, startElement, endElement)
	resp.Body.Close()
	return "", nil
}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}