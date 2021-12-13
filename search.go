package main

import (
    "fmt"
    "strings"
	"github.com/gocolly/colly/v2"
)

func main() {

    c := colly.NewCollector()

	c.OnHTML("table.memberdecls a.el", func(e *colly.HTMLElement) {
		link := e.Attr("href")

        if !(strings.HasPrefix(link, "group__") && strings.HasSuffix(link, ".html")) {
            return
        }

        fmt.Printf("%s\n%s\n", e.Text, strings.Repeat("-", len(e.Text)))
		c.Visit(e.Request.AbsoluteURL(link))
        fmt.Println()
	})

    c.OnHTML("table.memberdecls", func(e *colly.HTMLElement) {
        tableTitle := e.ChildText("tbody > .heading h2")

        if tableTitle != "Functions" { return }

        fmt.Printf("(%s)\n", tableTitle)

        e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
            cname := el.Attr("class")
            if strings.HasPrefix(cname, "separator") || cname == "Heading" {
                return
            } else if strings.HasPrefix(cname, "memitem") {
                fmt.Printf("%s ", el.ChildText(".memItemLeft"))
                fmt.Println(el.ChildText(".memItemRight"))
            } else if strings.HasPrefix(cname, "memdesc") {
                fmt.Println(el.ChildText(".mdescRight"))
            }
        })

    })

	c.Visit("https://llvm.org/doxygen/group__LLVMC.html")

}
