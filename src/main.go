package main

import (
    "net/http"
    "os"
    "fmt"
    "log"
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
    selection := doc.Find(".repo-list").Find("a.v-align-middle")
    selection.Each(func(i int, s *goquery.Selection) {
        // For each item found, get the title
        title := s.Text()
        fmt.Printf("%s\n", title)
    })
}

