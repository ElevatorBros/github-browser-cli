package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "strings"
	"github.com/PuerkitoBio/goquery"
)

func main() {
    webPage := "http://github.com/search?o=desc&s=stars&type=Repositories&q=" + os.Args[1]
    resp, err := http.Get(webPage)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    //title := doc.Find(".repo-list").Text()
    //fmt.Println(title)
    // var repos []Repo
    selection := doc.Find(".repo-list-item")
    selection.Each(func(i int, s *goquery.Selection) {
        data := s.Find("div.d-flex.flex-wrap.text-small.color-fg-muted").Text()
        data = strings.ReplaceAll(data, "\n", " ")
        data = strings.Trim(data, "  ")
        for _, elem := range strings.Split(data, "  ") {
            if elem == "" {
                continue
            }
            fmt.Printf("%s\n", strings.Trim(elem, " "))
        }
    })
}

