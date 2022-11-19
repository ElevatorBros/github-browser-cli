package main

import (
	"fmt"
	"log"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

// main function
func Fuzzy(repos []Repo) {
	idx, err := fuzzyfinder.FindMulti(
		repos,
		func(i int) string { 
            return repos[i].Name 
        },
		fuzzyfinder.WithPreviewWindow(
            func (i, w, h int) string {
                if i == -1 {
                    return ""
                }
                return fmt.Sprintf(
                    "Owner: %s\nName: %s\nStars: %d",
                    repos[i].Name,
                    repos[i].Stars,
                    repos[i].Language,
                    repos[i].License,
                    repos[i].Last_Updated,
                )
     }))

	if err != nil {
        log.Fatalf("Error performing fuzzy find: %s", err.Error())
	}

	//this prints what you selected after you press enter
	fmt.Printf("selected: %v\n", idx)
}
