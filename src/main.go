package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"fmt"
	//"io"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getRepos(search string, sortOrder string, pages int) []Repo {
	var repos []Repo
	for i := 1; i <= pages; i++ {
		page := strconv.Itoa(i)
		webPage := "http://github.com/search?o=desc&s=" + sortOrder + "&type=Repositories&q=" + search + "&p=" + page
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
		selection := doc.Find(".repo-list-item")
		selection.Each(func(i int, s *goquery.Selection) {
			data := s.Find("div.d-flex.flex-wrap.text-small.color-fg-muted").Text()
			data = strings.ReplaceAll(data, "\n", " ")
			//data = strings.Trim(data, "  ")
			var re = regexp.MustCompile("[ ]{2,}")
			data = re.ReplaceAllString(data, "$1    $2")
			repos = append(repos, Repo{
				Name:        s.Find("a.v-align-middle").Text(),
				Description: s.Find("p.mb-1").Text(),
				Bar:         data,
			})
		})
	}

	return repos
}

func printHelp() {
	fmt.Println("Usage: ./src search [args]")
	fmt.Println("\t-s : sort by stars")
	fmt.Println("\t-r : sort by recently updated")
	fmt.Println("\t-p [number of pages] : number of pages to include in search")
	fmt.Println("\t-n : disable read me support")
	fmt.Println("\t-u : search for a specific user")
	fmt.Println("\t-h : print this help")
}

func main() {
	// ./src search [args]
	// -p int : pages
	// -s : sort by stars (default relevence)
	// -r : sort by recently updated
	// -h : help
	search := ""
	sortOrder := ""
	pages := 1
	read := true

	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-s" {
			sortOrder = "stars"
		} else if os.Args[i] == "-r" {
			sortOrder = "updated"
		} else if os.Args[i] == "-n" {
			read = false
		} else if os.Args[i] == "-p" {
			newPages, err := strconv.Atoi(os.Args[i+1])
			if err != nil {
				printHelp()
				os.Exit(0)
			}
			pages = newPages
			i++
		} else if os.Args[i] == "-u" {
			if i == len(os.Args)-1 {
				printHelp()
				os.Exit(1)
			}
			search += "user%3A" + os.Args[i+1] + " "
			i++
		} else if os.Args[i] == "-h" {
			printHelp()
			os.Exit(0)
		} else {
			search += os.Args[i] + " "
		}
	}

	search = strings.ReplaceAll(search, " ", "%20")
	repos := getRepos(search, sortOrder, pages)
	repo := Fuzzy(repos)
	fmt.Printf("%s\n", repo.Name)
	if read {
		cmd := fmt.Sprintf("echo '%s' | vim -", repo.GetReadMe())
		cmd1 := exec.Command("bash", "-c", cmd)
		cmd1.Stdin = os.Stdin
		cmd1.Stdout = os.Stdout
		cmd1.Stderr = os.Stderr
		err := cmd1.Run()

		if err != nil {
		}

		//fmt.Printf("\n\n%s\n", repo.GetReadMe())
	}

	var reply string
	fmt.Print("Would you like to clone this repo? [Y/n]: ")
	fmt.Scanln(&reply)
	fmt.Println()
	if lower(reply) == "y" {
		cmd := exec.Command("git", "clone", fmt.Sprintf("git@github.com:%s.git", repo.Name))
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Error cloning repo: %s", err.Error())
		}
	}
}
