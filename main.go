package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

type header struct {
	text string
	id   string
}

func main() {
	h, err := os.Open("test.src.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(h)
	if err != nil {
		log.Fatal(err)
	}

	hx := []header{}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h3" {
			for _, a := range n.Attr {
				if a.Key == "id" {
					fmt.Println(a.Val)
					hx = append(hx, header{text: n.FirstChild.Data, id: a.Val})
					break
				}
			}
			fmt.Printf("%+v\n", n)
			// fmt.Println(n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	fmt.Println(hx)
}
