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
    selection := doc.Find(".repo-list")
    selection.Each(func(i int, s *goquery.Selection) {
        // For each item found, get the title
        name := s.Find("a.v-align-middle")
        var arr []Repo
        name.Each(func(i int, s *goquery.Selection) {
            // For each item found, get the title
            repoNameUser := strings.Split(s.Text(), "/")
            arr = append(arr, Repo{repoNameUser[0], repoNameUser[1], "", ""})

        })
        // description := s.Find("p.mb-1").Text()
        // stars := strings.Trim(s.Find("div.mr-3 a.Link--muted").Text(), " ")
        bar := s.Find("div.d-flex.flex-wrap.text-small.color-fg-muted").Find("div.mr-3")
        bar.Each(func(i int, s *goquery.Selection) {
            
            fmt.Println(strings.ReplaceAll(strings.Trim(s.Text(), " "), "\n", ""))
        })
        // repos = append(repos, Repo{
        //     Name: name,
        //     Description: description,
        //     Stars: stars,
        //     Language: s.Find("span[item='programmingLanguage']").Text(),
        //
        // })
    })
}

