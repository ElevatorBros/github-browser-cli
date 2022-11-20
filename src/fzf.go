package main

import (
	"fmt"
	//"go/format"
	"log"
	"strings"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

// main function
func Fuzzy(repos []Repo) int {
	idx, err := fuzzyfinder.FindMulti(
		repos,
		func(i int) string {
			return repos[i].Name
		},
		fuzzyfinder.WithPreviewWindow(
			func(i, w, h int) string {
				if i == -1 {
					return ""
				}
				return fmt.Sprintf(
					"Name: %s\nDesc: %s\nInfo: %s\n",
					repos[i].Name,
					repos[i].Description,
					strings.ReplaceAll(repos[i].Bar, "    ", "\n"),
				)
			}))

	if err != nil {
		log.Fatalf("Error performing fuzzy find: %s", err.Error())
	}

	return idx[0]
}
