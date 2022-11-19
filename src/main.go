package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
    args := os.Args[1:]
    fmt.Println(args)

    url := "https://www.github.com/search?q=" + args[0]
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(string(body))
}
