package main

import (
	"fmt"
	//"go/format"
	"log"
	"strings"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

// main function
func Fuzzy(repos []Repo) Repo {
	idx, err := fuzzyfinder.Find(
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
					strings.ReplaceAll(strings.ReplaceAll(repos[i].Bar, "    ", "\n"), "issues\nneed help", "issues need help"),
				)
			}))

	if err != nil {
		log.Fatalf("Error performing fuzzy find: %s", err.Error())
	}

	return repos[idx]
}
